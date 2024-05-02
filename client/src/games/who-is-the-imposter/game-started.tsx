import React from "react";
import { GameControl } from "./game-controls";

interface GameStartedProps {
  message: string;
  isAdmin: boolean;
  resetGame: () => void;
  decideWinner: (impostorWon: boolean) => void;
}

export const GameStarted = ({
  message,
  isAdmin,
  resetGame,
  decideWinner,
}: GameStartedProps) => {
  return (
    <>
      <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
        <h1 className="text-lg font-bold mb-4">Quem é o Impostor?</h1>
        <p className="mb-4 text-center">{message}</p>
      </div>
      {isAdmin && (
        <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
          <GameControl resetGame={resetGame} decideWinner={decideWinner} />
        </div>
      )}
    </>
  );
};

