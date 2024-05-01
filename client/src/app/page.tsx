import Image, { StaticImageData } from "next/image";
import Link from "next/link";

import imposterImg from "../../public/images/who-is-the-imposter.jpg";

type IGame = {
  id: number;
  name: string;
  slug: string;
  image: StaticImageData;
};

const GameCard = ({ game }: { game: IGame }) => {
  return (
    <Link href={`/${game.slug}`}>
      <b className="relative block rounded-lg shadow-lg p-4 m-2 h-64 w-80 transition duration-300 ease-in-out overflow-hidden group">
        <div className="absolute inset-0 z-0">
          <Image
            src={game.image}
            alt={`${game.name} Game Image`}
            layout="fill"
            objectFit="cover"
            className="group-hover:opacity-50"
          />
        </div>
        <div className="absolute inset-0 flex items-end justify-center p-4">
          <h2 className="text-2xl font-bold text-white drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,0.8)]">
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
    slug: "who-is-the-imposter",
    image: imposterImg,
  },
  // Add more games as needed
];

export default function Home() {
  return (
    <main className="bg-gray-800 min-h-screen flex flex-col items-center justify-center p-12">
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {games.map((game) => (
          <GameCard key={game.id} game={game} />
        ))}
      </div>
    </main>
  );
}
