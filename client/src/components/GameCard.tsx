import Link from "next/link";

export type Game = {
  name: string;
  slug: string;
};

interface GameCardProps {
  game: Game;
}

export const GameCard = ({ game }: GameCardProps) => {
  return (
    <Link href={`/${game.slug}`}>
      <b className="rounded-lg shadow-md p-4 m-2 flex flex-col justify-between">
        <h2 className="text-xl font-semibold">{game.name}</h2>
      </b>
    </Link>
  );
};
