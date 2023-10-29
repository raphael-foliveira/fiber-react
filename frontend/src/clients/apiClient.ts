import { HttpError } from '../errors/HttpError';

const BASE_URL = import.meta.env.VITE_API_URL;
const defaultHeaders = {
  'Content-Type': 'application/json',
};

export const apiClient = {
  get: async (endpoint: string, config: RequestInit = {}) => {
    const headers = getConfigHeaders(config);
    return fetchWithConfig(endpoint, {
      method: 'GET',
      ...config,
      headers,
    });
  },

  post: async (endpoint: string, body: unknown, config: RequestInit = {}) => {
    const headers = getConfigHeaders(config);
    return fetchWithConfig(endpoint, {
      method: 'POST',
      ...config,
      headers,
      body: JSON.stringify(body),
    });
  },

  put: async (endpoint: string, body: unknown, config: RequestInit = {}) => {
    const headers = getConfigHeaders(config);
    return fetchWithConfig(endpoint, {
      method: 'PUT',
      ...config,
      headers,
      body: JSON.stringify(body),
    });
  },

  delete: async (endpoint: string, config: RequestInit = {}) => {
    const headers = getConfigHeaders(config);
    return fetchWithConfig(endpoint, {
      method: 'DELETE',
      ...config,
      headers,
    });
  },

  patch: async (endpoint: string, body: unknown, config: RequestInit = {}) => {
    const headers = getConfigHeaders(config);
    return fetchWithConfig(endpoint, {
      method: 'PATCH',
      ...config,
      headers,
      body: JSON.stringify(body),
    });
  },
};

async function fetchWithConfig(endpoint: string, config: RequestInit = {}) {
  const response = await fetch(BASE_URL + endpoint, config);
  if (!response.ok) {
    console.error(await response.json());
    throw new HttpError(response.statusText, response.status);
  }
  return response.json();
}

function getConfigHeaders(config: RequestInit) {
  if (!config.headers) {
    return defaultHeaders;
  }
  return { ...defaultHeaders, ...config.headers };
}
