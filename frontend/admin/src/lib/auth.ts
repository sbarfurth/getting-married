const TOKEN_KEY = 'getting_married_token';

export function getToken(): string | undefined {
  return localStorage.getItem(TOKEN_KEY) ?? undefined;
}

export function setToken(newToken: string) {
  localStorage.setItem(TOKEN_KEY, newToken);
}

export function clearToken() {
  localStorage.removeItem(TOKEN_KEY);
}
