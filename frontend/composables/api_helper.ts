import { $fetch } from 'ofetch';

export const $apiHelper = $fetch.create({
  baseURL: '/a/',
  headers: {
    'Content-Type': 'application/json',
  },
});
