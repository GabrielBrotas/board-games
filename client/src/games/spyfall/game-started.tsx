import React from "react";
import { GameControl } from "./game-controls";

interface GameStartedProps {
  isAdmin: boolean;
  resetGame: () => void;
  decideWinner: (spiesWon: boolean) => void;
  inGame: boolean;
  role: string;
  location: string;
  showSpiesNumber: () => void;
}

export const SPY_ROLE = "spy";

export const GameStarted = ({
  isAdmin,
  resetGame,
  decideWinner,
  inGame,
  role,
  location,
  showSpiesNumber
}: GameStartedProps) => {
  return (
    <>
      <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
        <h1 className="text-lg font-bold mb-4 text-center">Quem é o Spy?</h1>
        {inGame ? (
          role == SPY_ROLE ? (
            <div className="">
              <p className="mb-4 text-center">Você é o Spy!</p>
            </div>
          ) : (
            <div className="flex flex-col items-start">
              <p className="mb-4 text-center">Localização: {location}</p>
              <p className="mb-4 text-center">Função: {role}</p>
            </div>
          )
        ) : (
          <p className="mb-4 text-center">
            Esperando jogadores terminar a partida...
          </p>
        )}
      </div>
      {isAdmin && (
        <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
          <GameControl
            resetGame={resetGame}
            decideWinner={decideWinner}
            showSpiesNumber={showSpiesNumber}
          />
        </div>
      )}
    </>
  );
};
