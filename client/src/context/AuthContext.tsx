// context/AuthContext.js
import { createContext, useContext, useState, useEffect } from "react";
import Cookies from "js-cookie";
import { api } from "@/lib/api";

type User = {
  id: string;
  name: string;
};

interface AuthContextType {
  user: User | null;
  isAdmin: boolean;
  isLoading: boolean;
  createUserOrLogin: (name: string) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const isAdmin = user?.name === "gbrotas";

  useEffect(() => {
    const user = Cookies.get("user");
    if (user) {
      const userParsed = JSON.parse(user);
      setUser(userParsed);
    }
    setIsLoading(false);
  }, []);

  const createUserOrLogin = async (name: string) => {
    const user = await api.createUserOrLogin(name);
    Cookies.set("user", JSON.stringify(user), { expires: 7 });
    setUser(user);
  };

  const logout = () => {
    setIsLoading(true);
    Cookies.remove("user");
    setUser(null);
    setIsLoading(false);
  };

  return (
    <AuthContext.Provider
      value={{ user, createUserOrLogin, logout, isAdmin, isLoading }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
