// context/AuthContext.js
import { createContext, useContext, useState, useEffect } from "react";
import Cookies from "js-cookie";

interface AuthContextType {
  user: string | null;
  login: (username: string) => void;
  changeUsername: (username: string) => void;
  isAdmin: boolean;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const isAdmin = user === "gbrotas";

  useEffect(() => {
    const username = Cookies.get("username");
    if (username) {
      setUser(username);
    }
    setIsLoading(false);
  }, []);

  const login = (username: string) => {
    Cookies.set("username", username, { expires: 7 });
    setUser(username);
  };

  const changeUsername = (username: string) => {
    Cookies.set("username", username, { expires: 7 });
    setUser(username);
  };

  return (
    <AuthContext.Provider
      value={{ user, login, changeUsername, isAdmin, isLoading }}
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
