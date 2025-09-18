declare module "vue" {
  interface ComponentCustomProperties {
    $t: (key: string, values?: Record<string, any>) => string
  }
}