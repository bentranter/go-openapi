// Code generated by mkgetter.go. DO NOT EDIT.

package openapi

func (v *OpenAPI) Openapi() string {
	return v.openapi
}

func (v *OpenAPI) Info() *Info {
	if v.info == nil {
		return &Info{}
	}
	return v.info
}

func (v *OpenAPI) Servers() []*Server {
	return v.servers
}

func (v *OpenAPI) Paths() *Paths {
	if v.paths == nil {
		return &Paths{}
	}
	return v.paths
}

func (v *OpenAPI) Components() *Components {
	if v.components == nil {
		return &Components{}
	}
	return v.components
}

func (v *OpenAPI) Security() []*SecurityRequirement {
	return v.security
}

func (v *OpenAPI) Tags() []*Tag {
	return v.tags
}

func (v *OpenAPI) ExternalDocs() *ExternalDocumentation {
	if v.externalDocs == nil {
		return &ExternalDocumentation{}
	}
	return v.externalDocs
}

func (v *OpenAPI) Extension() map[string]interface{} {
	return v.extension
}

func (v *Info) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Info) Title() string {
	return v.title
}

func (v *Info) Description() string {
	return v.description
}

func (v *Info) TermsOfService() string {
	return v.termsOfService
}

func (v *Info) Contact() *Contact {
	if v.contact == nil {
		return &Contact{}
	}
	return v.contact
}

func (v *Info) License() *License {
	if v.license == nil {
		return &License{}
	}
	return v.license
}

func (v *Info) Version() string {
	return v.version
}

func (v *Info) Extension() map[string]interface{} {
	return v.extension
}

func (v *Contact) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Contact) Name() string {
	return v.name
}

func (v *Contact) Url() string {
	return v.url
}

func (v *Contact) Email() string {
	return v.email
}

func (v *Contact) Extension() map[string]interface{} {
	return v.extension
}

func (v *License) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *License) Name() string {
	return v.name
}

func (v *License) Url() string {
	return v.url
}

func (v *License) Extension() map[string]interface{} {
	return v.extension
}

func (v *Server) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Server) Url() string {
	return v.url
}

func (v *Server) Description() string {
	return v.description
}

func (v *Server) Variables() map[string]*ServerVariable {
	return v.variables
}

func (v *Server) Extension() map[string]interface{} {
	return v.extension
}

func (v *ServerVariable) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *ServerVariable) Enum() []string {
	return v.enum
}

func (v *ServerVariable) Default_() string {
	return v.default_
}

func (v *ServerVariable) Description() string {
	return v.description
}

func (v *ServerVariable) Extension() map[string]interface{} {
	return v.extension
}

func (v *Components) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Components) Schemas() map[string]*Schema {
	return v.schemas
}

func (v *Components) Responses() map[string]*Response {
	return v.responses
}

func (v *Components) Parameters() map[string]*Parameter {
	return v.parameters
}

func (v *Components) Examples() map[string]*Example {
	return v.examples
}

func (v *Components) RequestBodies() map[string]*RequestBody {
	return v.requestBodies
}

func (v *Components) Headers() map[string]*Header {
	return v.headers
}

func (v *Components) SecuritySchemes() map[string]*SecurityScheme {
	return v.securitySchemes
}

func (v *Components) Links() map[string]*Link {
	return v.links
}

func (v *Components) Callbacks() map[string]*Callback {
	return v.callbacks
}

func (v *Components) Extension() map[string]interface{} {
	return v.extension
}

func (v *Paths) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Paths) Paths() map[string]*PathItem {
	return v.paths
}

func (v *Paths) Extension() map[string]interface{} {
	return v.extension
}

func (v *PathItem) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *PathItem) Summary() string {
	return v.summary
}

func (v *PathItem) Description() string {
	return v.description
}

func (v *PathItem) Get() *Operation {
	if v.get == nil {
		return &Operation{}
	}
	return v.get
}

func (v *PathItem) Put() *Operation {
	if v.put == nil {
		return &Operation{}
	}
	return v.put
}

func (v *PathItem) Post() *Operation {
	if v.post == nil {
		return &Operation{}
	}
	return v.post
}

func (v *PathItem) Delete() *Operation {
	if v.delete == nil {
		return &Operation{}
	}
	return v.delete
}

func (v *PathItem) Options() *Operation {
	if v.options == nil {
		return &Operation{}
	}
	return v.options
}

func (v *PathItem) Head() *Operation {
	if v.head == nil {
		return &Operation{}
	}
	return v.head
}

func (v *PathItem) Patch() *Operation {
	if v.patch == nil {
		return &Operation{}
	}
	return v.patch
}

func (v *PathItem) Trace() *Operation {
	if v.trace == nil {
		return &Operation{}
	}
	return v.trace
}

func (v *PathItem) Servers() []*Server {
	return v.servers
}

func (v *PathItem) Parameters() []*Parameter {
	return v.parameters
}

func (v *PathItem) Extension() map[string]interface{} {
	return v.extension
}

func (v *Operation) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Operation) Tags() []string {
	return v.tags
}

func (v *Operation) Summary() string {
	return v.summary
}

func (v *Operation) Description() string {
	return v.description
}

func (v *Operation) ExternalDocs() *ExternalDocumentation {
	if v.externalDocs == nil {
		return &ExternalDocumentation{}
	}
	return v.externalDocs
}

func (v *Operation) OperationID() string {
	return v.operationID
}

func (v *Operation) Parameters() *Parameter {
	if v.parameters == nil {
		return &Parameter{}
	}
	return v.parameters
}

func (v *Operation) RequestBody() *RequestBody {
	if v.requestBody == nil {
		return &RequestBody{}
	}
	return v.requestBody
}

func (v *Operation) Responses() *Responses {
	if v.responses == nil {
		return &Responses{}
	}
	return v.responses
}

func (v *Operation) Callbacks() map[string]*Callback {
	return v.callbacks
}

func (v *Operation) Deprecated() bool {
	return v.deprecated
}

func (v *Operation) Security() []*SecurityRequirement {
	return v.security
}

func (v *Operation) Servers() []*Server {
	return v.servers
}

func (v *Operation) Extension() map[string]interface{} {
	return v.extension
}

func (v *ExternalDocumentation) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *ExternalDocumentation) Description() string {
	return v.description
}

func (v *ExternalDocumentation) Url() string {
	return v.url
}

func (v *ExternalDocumentation) Extension() map[string]interface{} {
	return v.extension
}

func (v *Parameter) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Parameter) Name() string {
	return v.name
}

func (v *Parameter) In() string {
	return v.in
}

func (v *Parameter) Description() string {
	return v.description
}

func (v *Parameter) Required() bool {
	return v.required
}

func (v *Parameter) Deprecated() bool {
	return v.deprecated
}

func (v *Parameter) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *Parameter) Style() string {
	return v.style
}

func (v *Parameter) Explode() bool {
	return v.explode
}

func (v *Parameter) AllowReserved() bool {
	return v.allowReserved
}

func (v *Parameter) Schema() *Schema {
	if v.schema == nil {
		return &Schema{}
	}
	return v.schema
}

func (v *Parameter) Example() interface{} {
	return v.example
}

func (v *Parameter) Examples() map[string]*Example {
	return v.examples
}

func (v *Parameter) Content() map[string]*MediaType {
	return v.content
}

func (v *Parameter) Extension() map[string]interface{} {
	return v.extension
}

func (v *Parameter) Reference() string {
	return v.reference
}

func (v *RequestBody) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *RequestBody) Description() string {
	return v.description
}

func (v *RequestBody) Content() map[string]*MediaType {
	return v.content
}

func (v *RequestBody) Required() bool {
	return v.required
}

func (v *RequestBody) Extension() map[string]interface{} {
	return v.extension
}

func (v *RequestBody) Reference() string {
	return v.reference
}

func (v *MediaType) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *MediaType) Schema() *Schema {
	if v.schema == nil {
		return &Schema{}
	}
	return v.schema
}

func (v *MediaType) Example() interface{} {
	return v.example
}

func (v *MediaType) Examples() map[string]*Example {
	return v.examples
}

func (v *MediaType) Encoding() map[string]*Encoding {
	return v.encoding
}

func (v *MediaType) Extension() map[string]interface{} {
	return v.extension
}

func (v *Encoding) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Encoding) ContentType() string {
	return v.contentType
}

func (v *Encoding) Headers() map[string]*Header {
	return v.headers
}

func (v *Encoding) Style() string {
	return v.style
}

func (v *Encoding) Explode() string {
	return v.explode
}

func (v *Encoding) AllowReserved() bool {
	return v.allowReserved
}

func (v *Encoding) Extension() map[string]interface{} {
	return v.extension
}

func (v *Responses) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Responses) Responses() map[string]*Response {
	return v.responses
}

func (v *Responses) Extension() map[string]interface{} {
	return v.extension
}

func (v *Response) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Response) Description() string {
	return v.description
}

func (v *Response) Headers() map[string]*Header {
	return v.headers
}

func (v *Response) Content() map[string]*MediaType {
	return v.content
}

func (v *Response) Links() map[string]*Link {
	return v.links
}

func (v *Response) Extension() map[string]interface{} {
	return v.extension
}

func (v *Response) Reference() string {
	return v.reference
}

func (v *Callback) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Callback) Callback() map[string]*PathItem {
	return v.callback
}

func (v *Callback) Extension() map[string]interface{} {
	return v.extension
}

func (v *Callback) Reference() string {
	return v.reference
}

func (v *Example) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Example) Summary() string {
	return v.summary
}

func (v *Example) Description() string {
	return v.description
}

func (v *Example) Value() interface{} {
	return v.value
}

func (v *Example) ExternalVale() string {
	return v.externalVale
}

func (v *Example) Extension() map[string]interface{} {
	return v.extension
}

func (v *Example) Reference() string {
	return v.reference
}

func (v *Link) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Link) Operationreference() string {
	return v.operationreference
}

func (v *Link) OperationID() string {
	return v.operationID
}

func (v *Link) Parameters() map[string]interface{} {
	return v.parameters
}

func (v *Link) RequestBody() interface{} {
	return v.requestBody
}

func (v *Link) Description() string {
	return v.description
}

func (v *Link) Server() *Server {
	if v.server == nil {
		return &Server{}
	}
	return v.server
}

func (v *Link) Extension() map[string]interface{} {
	return v.extension
}

func (v *Link) Reference() string {
	return v.reference
}

func (v *Header) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Header) Name() string {
	return v.name
}

func (v *Header) In() string {
	return v.in
}

func (v *Header) Description() string {
	return v.description
}

func (v *Header) Required() bool {
	return v.required
}

func (v *Header) Deprecated() bool {
	return v.deprecated
}

func (v *Header) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *Header) Style() string {
	return v.style
}

func (v *Header) Explode() bool {
	return v.explode
}

func (v *Header) AllowReserved() bool {
	return v.allowReserved
}

func (v *Header) Schema() *Schema {
	if v.schema == nil {
		return &Schema{}
	}
	return v.schema
}

func (v *Header) Example() interface{} {
	return v.example
}

func (v *Header) Examples() map[string]*Example {
	return v.examples
}

func (v *Header) Content() map[string]*MediaType {
	return v.content
}

func (v *Header) Extension() map[string]interface{} {
	return v.extension
}

func (v *Header) Reference() string {
	return v.reference
}

func (v *Tag) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Tag) Name() string {
	return v.name
}

func (v *Tag) Description() string {
	return v.description
}

func (v *Tag) ExternalDocs() *ExternalDocumentation {
	if v.externalDocs == nil {
		return &ExternalDocumentation{}
	}
	return v.externalDocs
}

func (v *Tag) Extension() map[string]interface{} {
	return v.extension
}

func (v *Schema) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Schema) Title() string {
	return v.title
}

func (v *Schema) MultipleOf() int {
	return v.multipleOf
}

func (v *Schema) Maximum() int {
	return v.maximum
}

func (v *Schema) ExclusiveMaximum() bool {
	return v.exclusiveMaximum
}

func (v *Schema) Minimum() int {
	return v.minimum
}

func (v *Schema) ExclusiveMinimum() bool {
	return v.exclusiveMinimum
}

func (v *Schema) MaxLength() int {
	return v.maxLength
}

func (v *Schema) MinLength() int {
	return v.minLength
}

func (v *Schema) Pattern() string {
	return v.pattern
}

func (v *Schema) MaxItems() int {
	return v.maxItems
}

func (v *Schema) MinItems() int {
	return v.minItems
}

func (v *Schema) MaxProperties() int {
	return v.maxProperties
}

func (v *Schema) MinProperties() int {
	return v.minProperties
}

func (v *Schema) Required() []string {
	return v.required
}

func (v *Schema) Enum() []string {
	return v.enum
}

func (v *Schema) Type_() string {
	return v.type_
}

func (v *Schema) AllOf() []*Schema {
	return v.allOf
}

func (v *Schema) OneOf() []*Schema {
	return v.oneOf
}

func (v *Schema) AnyOf() []*Schema {
	return v.anyOf
}

func (v *Schema) Not() *Schema {
	if v.not == nil {
		return &Schema{}
	}
	return v.not
}

func (v *Schema) Items() *Schema {
	if v.items == nil {
		return &Schema{}
	}
	return v.items
}

func (v *Schema) Properties() map[string]*Schema {
	return v.properties
}

func (v *Schema) AdditionalProperties() *Schema {
	if v.additionalProperties == nil {
		return &Schema{}
	}
	return v.additionalProperties
}

func (v *Schema) Description() string {
	return v.description
}

func (v *Schema) Format() string {
	return v.format
}

func (v *Schema) Default_() string {
	return v.default_
}

func (v *Schema) Nullable() bool {
	return v.nullable
}

func (v *Schema) Discriminator() *Discriminator {
	if v.discriminator == nil {
		return &Discriminator{}
	}
	return v.discriminator
}

func (v *Schema) ReadOnly() bool {
	return v.readOnly
}

func (v *Schema) WriteOnly() bool {
	return v.writeOnly
}

func (v *Schema) Xml() *XML {
	if v.xml == nil {
		return &XML{}
	}
	return v.xml
}

func (v *Schema) ExternalDocs() *ExternalDocumentation {
	if v.externalDocs == nil {
		return &ExternalDocumentation{}
	}
	return v.externalDocs
}

func (v *Schema) Example() interface{} {
	return v.example
}

func (v *Schema) Deprecated() bool {
	return v.deprecated
}

func (v *Schema) Extension() map[string]interface{} {
	return v.extension
}

func (v *Schema) Reference() string {
	return v.reference
}

func (v *Discriminator) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *Discriminator) PropertyName() string {
	return v.propertyName
}

func (v *Discriminator) Mapping() map[string]string {
	return v.mapping
}

func (v *XML) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *XML) Name() string {
	return v.name
}

func (v *XML) Namespace() string {
	return v.namespace
}

func (v *XML) Prefix() string {
	return v.prefix
}

func (v *XML) Attribute() bool {
	return v.attribute
}

func (v *XML) Wrapped() bool {
	return v.wrapped
}

func (v *XML) Extension() map[string]interface{} {
	return v.extension
}

func (v *SecurityScheme) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *SecurityScheme) Type_() string {
	return v.type_
}

func (v *SecurityScheme) Description() string {
	return v.description
}

func (v *SecurityScheme) Name() string {
	return v.name
}

func (v *SecurityScheme) In() string {
	return v.in
}

func (v *SecurityScheme) Scheme() string {
	return v.scheme
}

func (v *SecurityScheme) BearerFormat() string {
	return v.bearerFormat
}

func (v *SecurityScheme) Flows() *OAuthFlows {
	if v.flows == nil {
		return &OAuthFlows{}
	}
	return v.flows
}

func (v *SecurityScheme) OpenIDConnectURL() string {
	return v.openIDConnectURL
}

func (v *SecurityScheme) Extension() map[string]interface{} {
	return v.extension
}

func (v *SecurityScheme) Reference() string {
	return v.reference
}

func (v *OAuthFlows) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *OAuthFlows) Implicit() *OAuthFlow {
	if v.implicit == nil {
		return &OAuthFlow{}
	}
	return v.implicit
}

func (v *OAuthFlows) Password() *OAuthFlow {
	if v.password == nil {
		return &OAuthFlow{}
	}
	return v.password
}

func (v *OAuthFlows) ClientCredentials() *OAuthFlow {
	if v.clientCredentials == nil {
		return &OAuthFlow{}
	}
	return v.clientCredentials
}

func (v *OAuthFlows) AuthorizationCode() *OAuthFlow {
	if v.authorizationCode == nil {
		return &OAuthFlow{}
	}
	return v.authorizationCode
}

func (v *OAuthFlows) Extension() map[string]interface{} {
	return v.extension
}

func (v *OAuthFlow) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *OAuthFlow) AuthorizationURL() string {
	return v.authorizationURL
}

func (v *OAuthFlow) TokenURL() string {
	return v.tokenURL
}

func (v *OAuthFlow) RefreshURL() string {
	return v.refreshURL
}

func (v *OAuthFlow) Scopes() map[string]string {
	return v.scopes
}

func (v *OAuthFlow) Extension() map[string]interface{} {
	return v.extension
}

func (v *SecurityRequirement) Root() *OpenAPI {
	if v.root == nil {
		return &OpenAPI{}
	}
	return v.root
}

func (v *SecurityRequirement) SecurityRequirement() map[string][]string {
	return v.securityRequirement
}