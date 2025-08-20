variable "zone_id" {
  description = "The zone ID to configure"
  type        = string
}

variable "automatic_https_rewrites" {
  description = "Enable automatic HTTPS rewrites"
  type        = string
  default     = "off"
}

variable "min_tls_version" {
  description = "Minimum TLS version"
  type        = string
  default     = "1.2"
}

variable "ssl" {
  description = "SSL setting"
  type        = string
  default     = "flexible"
}

variable "ipv6" {
  description = "Enable IPv6"
  type        = string
  default     = "off"
}

variable "websockets" {
  description = "Enable WebSockets"
  type        = string
  default     = "off"
}

variable "browser_check" {
  description = "Browser integrity check"
  type        = string
  default     = "on"
}

variable "security_level" {
  description = "Security level"
  type        = string
  default     = "medium"
}

variable "email_obfuscation" {
  description = "Email obfuscation"
  type        = string
  default     = "on"
}

variable "always_use_https" {
  description = "Always use HTTPS"
  type        = string
  default     = "off"
}

variable "security_header_enabled" {
  description = "Enable security headers"
  type        = bool
  default     = false
}

variable "security_header_include_subdomains" {
  description = "Include subdomains in security headers"
  type        = bool
  default     = false
}

variable "security_header_max_age" {
  description = "Max age for security headers"
  type        = number
  default     = 86400
}

variable "security_header_nosniff" {
  description = "Enable nosniff security header"
  type        = bool
  default     = false
}

variable "security_header_preload" {
  description = "Enable preload security header"
  type        = bool
  default     = false
}

variable "enable_network_error_logging" {
  description = "Enable network error logging"
  type        = bool
  default     = false
}