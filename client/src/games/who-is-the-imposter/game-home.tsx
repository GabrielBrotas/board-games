import React from "react";
import { IPlayer } from ".";
import { MdDelete } from "react-icons/md";
interface GameHomeProps {
  players: IPlayer[];
  startGame: () => void;
  isAdmin: boolean;
  removePlayer: (id: string) => void;
}

export const GameHome = ({
  players,
  startGame,
  isAdmin,
  removePlayer,
}: GameHomeProps) => {
  return (
    <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg">
      <h1 className="text-lg font-bold mb-4">Jogadores na sala:</h1>
      <ul className="mb-4">
        {players.map((player) => (
          <li key={player.name} className="bg-gray-700 p-2 rounded-md mb-2">
            <div className="flex justify-between items-center">
              <p>{player.name}</p>
              <div className="flex justify-between items-center gap-4">
                <span>{player.points}</span>
                {isAdmin && (
                  <button
                    className="text-red-500"
                    title="Remover jogador"
                    onClick={() => removePlayer(player.id)}
                  >
                    <MdDelete color="#ff0000" size={20} />
                  </button>
                )}
              </div>
            </div>
          </li>
        ))}
      </ul>

      {isAdmin ? (
        <button
          className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
          onClick={startGame}
        >
          Start Game
        </button>
      ) : (
        <p>Esperando o host iniciar o jogo</p>
      )}
    </div>
  );
};
