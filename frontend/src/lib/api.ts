import axios from "axios";
import { AuthResponse } from "./validations/auth";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000/api";

// Create an axios instance
export const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add a request interceptor to add authorization header
api.interceptors.request.use(
  (config) => {
    // Get token from localStorage if we're on the client
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("auth_token");
      if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
      }
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export const authService = {
  async signup(userData: {
    email: string;
    password: string;
    firstName?: string;
    lastName?: string;
  }): Promise<AuthResponse> {
    const response = await api.post<AuthResponse>("/auth/signup", userData);

    // Save token in localStorage
    if (typeof window !== "undefined") {
      localStorage.setItem("auth_token", response.data.accessToken);
    }

    return response.data;
  },

  async login(credentials: {
    email: string;
    password: string;
  }): Promise<AuthResponse> {
    const response = await api.post<AuthResponse>("/auth/login", credentials);

    // Save token in localStorage
    if (typeof window !== "undefined") {
      localStorage.setItem("auth_token", response.data.accessToken);
    }

    return response.data;
  },

  async getCurrentUser(): Promise<AuthResponse["user"] | null> {
    try {
      const response = await api.get<{ user: AuthResponse["user"] }>(
        "/auth/me"
      );
      return response.data.user;
    } catch (error) {
      return null;
    }
  },

  async logout(): Promise<void> {
    if (typeof window !== "undefined") {
      localStorage.removeItem("auth_token");
    }
  },
};
