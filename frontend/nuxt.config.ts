// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  extends: [
    // '@nuxt/ui-pro',
  ],

  modules: ['@nuxt/ui', '@nuxtjs/i18n'],

  i18n: {
    vueI18n: './i18n.config.ts',
  },

  components: [
    {
      path: '~/components',
      pathPrefix: false,
    },
  ],

  ui: {
    // icons: ["material-symbols"]
  },

  nitro: {
    devProxy: {
      '/a/': 'http://localhost:8080/a/',
    },
  },

  compatibilityDate: '2024-08-08',
});