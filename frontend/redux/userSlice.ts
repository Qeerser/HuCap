import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";

interface User {
  email: string;
  username?: string;
  password: string;
  full_name?: string;
}

interface UserState {
  user: User | null;
  loading: boolean;
  error: string | null;
  token: string | null;
}

// Initial state
const initialState: UserState = {
  user: null,
  loading: false,
  error: null,
  token: null,
};

// Async thunk for user registration
export const registerUser = createAsyncThunk(
  "user/register",
  async (userData: User, { rejectWithValue }) => {
    try {
      const response = await axios.post("http://localhost:5000/api/register", userData, {
        headers: { "Content-Type": "application/json" },
      });
      return response.data;
    } catch (error: any) {
      return rejectWithValue(error.response?.data || "Registration failed");
    }
  }
);

// Async thunk for user login
export const loginUser = createAsyncThunk(
  "user/login",
  async (userData: Pick<User, "email" | "password">, { rejectWithValue }) => {
    try {
      const response = await axios.post("http://localhost:5000/api/login", userData, {
        headers: { "Content-Type": "application/json" },
      });
      return response.data;
    } catch (error: any) {
      return rejectWithValue(error.response?.data || "Login failed");
    }
  }
);

// User slice
const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    logout: (state) => {
      state.user = null;
      state.token = null;
      state.error = null;
    },
  },
  extraReducers: (builder) => {
    builder
      // Handle Register
      .addCase(registerUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(registerUser.fulfilled, (state, action) => {
        state.loading = false;
        state.user = action.payload;
      })
      .addCase(registerUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })

      // Handle Login
      .addCase(loginUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.loading = false;
        state.user = action.payload.user;
        state.token = action.payload.token;
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.loading = false;
        state.error = typeof action.payload === "string" ? action.payload : "Login failed";
      });
  },
});

export const { logout } = userSlice.actions;
export default userSlice.reducer;
