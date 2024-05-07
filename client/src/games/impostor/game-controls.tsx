import React from "react";

interface GameControlProps {
  resetGame: () => void;
  decideWinner: (impostorWon: boolean) => void;
  showImpostorsNumber: () => void;
}

export const GameControl = ({
  resetGame,
  decideWinner,
  showImpostorsNumber,
}: GameControlProps) => {
  return (
    <div className="flex flex-col items-center justify-center w-full min-w-44 gap-4">
      <p className="text-xl font-medium">Vencedor</p>
      <button
        onClick={() => decideWinner(true)}
        className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
      >
        Impostor
      </button>
      <button
        onClick={() => decideWinner(false)}
        className="bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
      >
        Civilians
      </button>
      <button
        className="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
        onClick={showImpostorsNumber}
      >
        Show How Many Impostors
      </button>
      <button
        className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
        onClick={resetGame}
      >
        Reset Game
      </button>
    </div>
  );
};
