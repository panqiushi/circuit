// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  extends: [
    // '@nuxt/ui-pro',
  ],
  modules: ['@nuxt/ui', '@nuxtjs/i18n', "@sidebase/nuxt-auth"],
  auth: {
    baseURL: "/a",
    provider: {
      type: 'local',
      endpoints: {
        getSession: false,
        signIn: { path: '/login', method: 'post' },
      },
      pages: {
        login: '/dashboard'
      },
      token: {
        signInResponseTokenPointer: '/auth/token',
      },
      sessionDataType: { id: 'string', email: 'string', name: 'string', role: 'admin | guest | account', subscriptions: "{ id: number, status: 'ACTIVE' | 'INACTIVE' }[]" },
    },
    session: {
      // Whether to refresh the session every time the browser window is refocused.
      enableRefreshOnWindowFocus: false,

      // Whether to refresh the session every `X` milliseconds. Set this to `false` to turn it off. The session will only be refreshed if a session already exists.
      enableRefreshPeriodically: false
    },
    globalAppMiddleware: {
      isEnabled: true
    }
  },
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
});