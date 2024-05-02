"use client";
import Image, { StaticImageData } from "next/image";
import Link from "next/link";

import imposterImg from "../../public/images/who-is-the-imposter.jpg";
import { useState, useEffect } from "react";
import { useAuth } from "@/context/AuthContext";

type IGame = {
  id: number;
  name: string;
  slug: string;
  image: StaticImageData;
};

const GameCard = ({ game }: { game: IGame }) => {
  return (
    <Link href={`/${game.slug}`}>
      <b className="relative block rounded-lg shadow-lg h-full w-full sm:w-96 transition duration-300 ease-in-out overflow-hidden group ">
        <div className="inset-0 z-0">
          <Image
            src={game.image}
            alt={`${game.name} Game Image`}
            objectFit="cover"
            className="group-hover:opacity-50"
          />
        </div>
        <div className="absolute inset-0 z-10 flex items-end justify-center p-4">
          <h2 className="text-xl sm:text-2xl font-bold text-white drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,0.9)]">
            {game.name}
          </h2>
        </div>
        <div className="absolute inset-x-0 bottom-0 h-0 bg-gradient-to-t from-black to-transparent transition-all duration-300 ease-in-out group-hover:h-1/2"></div>
      </b>
    </Link>
  );
};

const UserInfo = () => {
  const { user, login } = useAuth();

  const [newName, setNewName] = useState("");
  const [isEditing, setIsEditing] = useState(false);
  const [showSuccess, setShowSuccess] = useState(false);

  useEffect(() => {
    if (user) {
      setNewName(user);
    }
  }, [user]);

  const handleSaveName = () => {
    if (!newName) return;
    login(newName); // Assuming 'login' function also updates the username
    setIsEditing(false);
    setShowSuccess(true);
    setTimeout(() => setShowSuccess(false), 3000); // Hide notification after 3 seconds
  };

  const handleEdit = () => {
    setIsEditing(true);
  };

  const handleCancel = () => {
    setIsEditing(false);
    setNewName(user as string);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.value.length > 15) return;

    setNewName(e.target.value);
  };

  return (
    <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
      {isEditing || !user ? (
        <div className="flex flex-col flex-1 items-center justify-center">
          <input
            type="text"
            value={newName}
            onChange={handleChange}
            placeholder="Username"
            className="bg-gray-700 text-white mb-4 p-2 rounded w-full"
          />
          <button
            onClick={handleSaveName}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full mb-2"
          >
            {!user ? "Login" : "Atualizar Nome"}
          </button>
          {user && (
            <button
              onClick={handleCancel}
              className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
            >
              Cancelar
            </button>
          )}
        </div>
      ) : (
        <div className="flex flex-col flex-1 items-center justify-center gap-4">
          <p>{user}</p>
          <button
            onClick={handleEdit}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full mb-4"
          >
            Atualizar Nome
          </button>
        </div>
      )}
      {showSuccess && (
        <div className="text-green-500 p-2 rounded text-center">
          Nome atualizado com sucesso!
        </div>
      )}
    </div>
  );
};

const games = [
  {
    id: 1,
    name: "Quem é o Impostor?",
    slug: "who-is-the-imposter",
    image: imposterImg,
  },
];

export default function Home() {
  const { user, isLoading } = useAuth();

  if (isLoading) {
    return (
      <div className="bg-gray-800 min-h-screen flex flex-col items-center p-12 justify-center">
        <div role="status">
          <svg
            aria-hidden="true"
            className="w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
            viewBox="0 0 100 101"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
              fill="currentColor"
            />
            <path
              d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
              fill="currentFill"
            />
          </svg>
          <span className="sr-only">Loading...</span>
        </div>
      </div>
    );
  }
  return (
    <main className="bg-gray-800 min-h-screen flex flex-col items-center p-12">
      <div>
        <UserInfo />
      </div>

      <div className="container flex flex-col mt-4">
        <h1 className="text-4xl text-center font-bold text-white mb-8">
          Games
        </h1>

        {user ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            {games.map((game) => (
              <GameCard key={game.id} game={game} />
            ))}
          </div>
        ) : (
          <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm mx-auto">
            <p className="mb-4 text-center items-center">
              Você precisa de uma username para poder jogar!
            </p>
          </div>
        )}
      </div>
    </main>
  );
}
