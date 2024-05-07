"use client";
import Image, { StaticImageData } from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { toast } from "react-hot-toast";

import imposterImg from "../../../public/images/who-is-the-imposter.jpg";
import spyfallImg from "../../../public/images/spyfall.jpg";
import { useAuth } from "@/context/AuthContext";
import { Loading } from "@/components/loading";

type IGame = {
  id: number;
  name: string;
  slug: string;
  image: StaticImageData;
};

const GameCard = ({ game }: { game: IGame }) => {
  return (
    <Link href={`/games/${game.slug}`}>
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

const games = [
  {
    id: 1,
    name: "Quem é o Impostor?",
    slug: "imposter",
    image: imposterImg,
  },
  {
    id: 2,
    name: "Spyfall",
    slug: "spyfall",
    image: spyfallImg,
  }
];

export default function Home() {
  const router = useRouter();
  const { user, isLoading, logout } = useAuth();

  const handleLogout = () => {
    logout();
    router.push("/");
    toast.success("Logout realizado com sucesso");
  };

  if (isLoading) {
    return <Loading />;
  }

  return (
    <main className="bg-gray-800 min-h-screen flex flex-col items-center p-12">
      <div>
        <button
          onClick={handleLogout}
          className="absolute top-0 left-0 m-4 bg-gray-500 text-white px-4 py-2 rounded"
        >
          Sair
        </button>
      </div>

      <div className="container flex flex-col mt-4">
        <h1 className="text-4xl text-center font-bold text-white mb-8">
          Games
        </h1>

        {user ? (
          // <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          <div className="flex gap-4 flex-wrap justify-center items-center">
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
