{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "labels.common" -}}
{{ include "labels.selector" . }}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
app.kubernetes.io/name: {{ .Release.Name | quote }}
app.kubernetes.io/instance: {{ .Release.Name | quote }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
giantswarm.io/service-type: "{{ .Values.serviceType }}"
helm.sh/chart: {{ include "chart" . | quote }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "labels.selector" -}}
app: {{ .Release.Name | quote }}
{{- end -}}

{{/*
Create the list of domains to update
*/}}
{{- define "domain.list" -}}
{{- if .Values.externalDNS.domainFilterList -}}
{{- range .Values.externalDNS.domainFilterList }}
{{ printf "- --domain-filter=%s" . }}
{{- end }}
{{- else }}
{{- if eq .Values.aws.access "external" }}
{{- printf "- --domain-filter=%s" .Values.aws.baseDomain }}
{{- else }}
{{- printf "- --domain-filter=%s" .Values.baseDomain }}
{{- end }}
{{- end }}
{{- end }}
