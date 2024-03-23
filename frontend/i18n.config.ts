export default defineI18nConfig(() => ({
  legacy: false,
  locale: 'zh',
  messages: {
    en: {
      welcome: 'Welcome',
      login: 'Login',
      signup: 'Signup',
      emial: 'Email',
      password: 'Password',
      confirm: 'Confirm',
      submit: 'Submit',
    },
    zh: {
      welcome: '欢迎',
      login: '登录',
      signup: '注册',
      emial: '邮箱',
      password: '密码',
      confirm: '确认',
      submit: '提交',
    },
  },
}));
