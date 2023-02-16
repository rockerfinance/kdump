package kubectl

import (
	"github.com/gigurra/go-util/shell"
	"github.com/gigurra/go-util/stringutil"
	"github.com/gigurra/kdump/internal/k8s"
	"github.com/samber/lo"
	"log"
	"os/exec"
	"strings"
)

func init() {

	// Check we have
	log.Printf("Checking that kubectl is installed...")
	if !shell.CommandExists("kubectl") {
		panic("kubectl not on path!")
	}

	log.Printf("Checking that kubectl neat is installed...")
	shell.RunCommand("kubectl", "neat", "--help")
}

func Namespaces() []string {
	return stringutil.MapStrArray(stringutil.SplitLines(runCommand("get", "namespaces", "-o", "name")), removeK8sResourcePrefix)
}

func CurrentNamespace() string {
	return runCommand("config", "view", "--minify", "--output", "jsonpath={..namespace}")
}

func CurrentContext() string {
	return runCommand("config", "current-context")
}

func ListNamespacedResourcesOfType(namespace string, resourceType string) []string {
	rawString := runCommand("-n", namespace, "get", resourceType, "-o", "name")
	itemsWithPrefixes := stringutil.RemoveEmptyLines(stringutil.SplitLines(rawString))
	return lo.Map(itemsWithPrefixes, func(in string, _ int) string {
		return stringutil.RemoveUpToAndIncluding(in, "/")
	})
}

func DownloadNamespacedResource(namespace string, resourceType string, resourceName string, format string) string {
	rawString := runCommand("-n", namespace, "get", resourceType, resourceName, "-o", format)
	return rawString
}

func DownloadGlobalResource(resourceType string, resourceName string, format string) string {
	rawString := runCommand("get", resourceType, resourceName, "-o", format)
	return rawString
}

func DownloadEverything(types []*k8s.ApiResourceType) string {

	qualifiedTypeNames := lo.Map(types, func(in *k8s.ApiResourceType, _ int) string {
		return in.QualifiedName
	})

	return runCommand("get", strings.Join(qualifiedTypeNames, ","), "--all-namespaces", "-o", "yaml")
}

func DownloadEverythingInNamespace(types []*k8s.ApiResourceType, namespace string) string {

	qualifiedTypeNames := lo.Map(types, func(in *k8s.ApiResourceType, _ int) string {
		return in.QualifiedName
	})

	return runCommand("get", strings.Join(qualifiedTypeNames, ","), "-n", namespace, "-o", "yaml")
}

func DownloadEverythingOfType(tpe *k8s.ApiResourceType) string {
	return runCommand("get", tpe.QualifiedName, "-o", "yaml")
}

func DownloadEverythingOfTypeInNamespace(tpe *k8s.ApiResourceType, namespace string) string {
	return runCommand("get", "-n", namespace, tpe.QualifiedName, "-o", "yaml")
}

type ApiResourceTypesResponse struct {
	All        []*k8s.ApiResourceType
	Accessible ApiResourceTypesAccessible
}

type ApiResourceTypesAccessible struct {
	All        []*k8s.ApiResourceType
	Global     []*k8s.ApiResourceType
	Namespaced []*k8s.ApiResourceType
}

func ApiResourceTypes() ApiResourceTypesResponse {

	rawString := runCommand("api-resources", "-o", "wide")

	_ /* schema */, apiResourcesRaw := stringutil.ParseStdOutTable(rawString)

	allApiResources := lo.Map(apiResourcesRaw, func(in map[string]string, _ int) *k8s.ApiResourceType {

		out := &k8s.ApiResourceType{
			Name:       stringutil.MapStrValOrElse(in, "NAME", ""),
			ShortNames: stringutil.CsvStr2arr(stringutil.MapStrValOrElse(in, "SHORTNAMES", "")),
			Namespaced: stringutil.Str2boolOrElse(stringutil.MapStrValOrElse(in, "NAMESPACED", ""), false),
			Kind:       stringutil.MapStrValOrElse(in, "KIND", ""),
			Verbs:      stringutil.WierdKubectlArray2arr(stringutil.MapStrValOrElse(in, "VERBS", "")),
		}

		apiVersionString := stringutil.MapStrValOrElse(in, "APIVERSION", "")
		parts := strings.Split(apiVersionString, "/")
		if len(parts) == 1 {
			out.ApiVersion = k8s.ApiVersion{Version: parts[0]}
			out.QualifiedName = out.Name
		} else if len(parts) == 2 {
			out.ApiVersion = k8s.ApiVersion{Version: parts[1], Name: parts[0]}
			out.QualifiedName = out.Name + "." + out.ApiVersion.Name
		} else {
			panic("Unable to parse ApiVersion from string: " + apiVersionString)
		}

		return out

	})

	accessibleApiResources := lo.Filter(allApiResources, func(r *k8s.ApiResourceType, _ int) bool { return lo.Contains(r.Verbs, "get") })
	globalResources := lo.Filter(accessibleApiResources, func(r *k8s.ApiResourceType, _ int) bool { return !r.Namespaced })
	namespacedResources := lo.Filter(accessibleApiResources, func(r *k8s.ApiResourceType, _ int) bool { return r.Namespaced })

	return ApiResourceTypesResponse{
		All: allApiResources,
		Accessible: ApiResourceTypesAccessible{
			All:        accessibleApiResources,
			Global:     globalResources,
			Namespaced: namespacedResources,
		},
	}
}

func runCommand(args ...string) string {

	fullCommand := "kubectl " + strings.Join(args, " ")

	cmd := exec.Command("kubectl", args...)

	outputBytes, err := cmd.Output()

	if err != nil {
		panic(`command "` + fullCommand + `" failed with error: ` + err.Error())
	}

	return strings.TrimSpace(string(outputBytes))
}

func removeK8sResourcePrefix(in string) string {
	return stringutil.RemoveUpToAndIncluding(in, "/")
}