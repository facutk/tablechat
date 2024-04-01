export const API_URL =
  import.meta.env.MODE === 'production'
    ? 'https://api.tablechat.me'
    : 'http://localhost:3000';
