"use client";
import { useState, useEffect } from "react";
import { useAuth } from "@/context/AuthContext";
import { toast } from "react-hot-toast";
import { useRouter } from "next/navigation";
import { Loading } from "@/components/loading";

export default function Home() {
  const { isLoading, createUserOrLogin } = useAuth();
  const router = useRouter();

  const [name, setName] = useState("");

  const handleLogin = async () => {
    if (!name) {
      toast.error("Nome é obrigatório");
      return;
    }
    if (name.length > 15) {
      toast.error("Nome deve ser menor que 15 caracteres");
      return;
    }

    try {
      await createUserOrLogin(name);
      toast.success("Login realizado com sucesso");
      router.push("/games");
    } catch (error) {
      toast.error(`Error no login: ${(error as any)?.response?.data}`);
    }
  };

  if (isLoading) {
    return <Loading />;
  }

  return (
    <main className="bg-gray-800 min-h-screen flex flex-col justify-center items-center p-12">
      <div>
        <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
          <div className="flex flex-col flex-1 items-center justify-center gap-2">
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Username"
              className="bg-gray-700 text-white mb-4 p-2 rounded w-full"
            />
            <button
              onClick={handleLogin}
              className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full mb-2"
            >
              Login
            </button>
          </div>
        </div>
      </div>
    </main>
  );
}
