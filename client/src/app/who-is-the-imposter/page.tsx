"use client";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function WhoIsTheImposter() {
  const [players, setPlayers] = useState<string[]>([]);
  const [message, setMessage] = useState("");
  const [gameStarted, setGameStarted] = useState(false);

  const [imposterChances, setImposterChances] = useState({
    one: 100,
    two: 0,
    three: 0,
  });
  const [category, setCategory] = useState("");

  const [currentName, setCurrentName] = useState("");
  const isAdmin = currentName == "gbrotas";

  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const newWs = new WebSocket("ws://localhost:8081/ws");
    setWs(newWs);

    newWs.onopen = () => console.log("Connected to server");

    newWs.onmessage = function (event) {
      const data = JSON.parse(event.data);
      if (data.type === "playerList") {
        setPlayers(data.players);
      } else if (data.type === "role") {
        setMessage(data.wordOrRole);
        setGameStarted(true);
      } else if (data.type === "resetGame") {
        setGameStarted(false);
        setMessage("");
      }
    };

    newWs.onclose = function () {
      console.log("Connection closed");
    };

    return () => {
      console.log("Closing connection...");
      newWs.close();
    };
  }, []);

  const handleNameChange = (name: string) => {
    if (!ws) {
      console.error("WebSocket is not initialized.");
      return;
    }
    const nameFormated = name.trim();

    if (!nameFormated) return;

    if (ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "changeName", newName: nameFormated }));
      setCurrentName(nameFormated);
    } else {
      console.error("WebSocket is not open.");
    }
  };

  const startGame = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(
        JSON.stringify({ type: "startGame", category, ...imposterChances })
      );
    } else {
      console.error("WebSocket is not open.");
    }
  };

  const resetGame = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "resetGame" }));
    } else {
      console.error("WebSocket is not open.");
    }
  };

  const handleImposterChange = (key: string, value: string) => {
    const newChances = { ...imposterChances, [key]: parseInt(value) };
    setImposterChances(newChances);
  };

  return (
    <div className="bg-gray-800 min-h-screen flex flex-col items-center justify-center gap-4">
      <Link href="/">
        <b className="absolute top-0 left-0 m-4 bg-gray-500 text-white px-4 py-2 rounded">
          Voltar
        </b>
      </Link>
      {!gameStarted ? (
        <GameHome
          players={players}
          handleNameChange={handleNameChange}
          startGame={startGame}
          isAdmin={isAdmin}
        />
      ) : (
        <GameStarted
          message={message}
          isAdmin={isAdmin}
          resetGame={resetGame}
        />
      )}

      {isAdmin && !gameStarted && (
        <GameSetup
          imposterChances={imposterChances}
          handleImposterChange={handleImposterChange}
          category={category}
          setCategory={setCategory}
        />
      )}
    </div>
  );
}

interface GameHomeProps {
  players: string[];
  handleNameChange: (name: string) => void;
  startGame: () => void;
  isAdmin: boolean;
}

const GameHome = ({
  players,
  isAdmin,
  startGame,
  handleNameChange,
}: GameHomeProps) => {
  const [newName, setNewName] = useState("");

  return (
    <>
      <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
        <h1 className="text-lg font-bold mb-4">Jogadores na sala:</h1>
        <ul className="mb-4">
          {players.map((player, index) => (
            <li key={index} className="bg-gray-700 p-2 rounded-md mb-2">
              {player}
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

      <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
        <input
          type="text"
          value={newName}
          onChange={(e) => setNewName(e.target.value)}
          placeholder="New name"
          className="bg-gray-700 text-white mb-4 p-2 rounded w-full"
        />
        <button
          onClick={() => {
            handleNameChange(newName);
            setNewName("");
          }}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full mb-4"
        >
          Mudar Nome
        </button>
      </div>
    </>
  );
};

interface GameStartedProps {
  message: string;
  isAdmin: boolean;
  resetGame: () => void;
}

const GameStarted = ({ isAdmin, message, resetGame }: GameStartedProps) => (
  <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
    <h1 className="text-lg font-bold mb-4">Quem é o Impostor?</h1>
    <p className="mb-4 text-center text-">{message}</p>
    {isAdmin && (
      <button
        className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
        onClick={resetGame}
      >
        Reset Game
      </button>
    )}
  </div>
);

interface GameSetupProps {
  imposterChances: {
    one: number;
    two: number;
    three: number;
  };
  handleImposterChange: (key: string, value: string) => void;
  category: string;
  setCategory: (category: string) => void;
}

const GameSetup = ({
  handleImposterChange,
  imposterChances,
  category,
  setCategory,
}: GameSetupProps) => (
  <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-lg w-full">
    <h1 className="text-xl font-bold mb-6 text-center">Setup Game</h1>
    <div className="flex justify-between items-center mb-4">
      <label htmlFor="one-imposter" className="flex-1 text-sm font-semibold">
        1 Imposter Chance:
      </label>
      <input
        type="range"
        id="one-imposter"
        value={imposterChances.one}
        onChange={(e) => handleImposterChange("one", e.target.value)}
        className="flex-2 range range-primary"
        min="0"
        max="100"
      />
      <span className="ml-4 w-12 text-center">{imposterChances.one}%</span>
    </div>
    <div className="flex justify-between items-center mb-4">
      <label htmlFor="two-imposters" className="flex-1 text-sm font-semibold">
        2 Imposters Chance:
      </label>
      <input
        type="range"
        id="two-imposters"
        value={imposterChances.two}
        onChange={(e) => handleImposterChange("two", e.target.value)}
        className="flex-2 range range-primary"
        min="0"
        max="100"
      />
      <span className="ml-4 w-12 text-center">{imposterChances.two}%</span>
    </div>
    <div className="flex justify-between items-center mb-6">
      <label htmlFor="three-imposters" className="flex-1 text-sm font-semibold">
        3 Imposters Chance:
      </label>
      <input
        type="range"
        id="three-imposters"
        value={imposterChances.three}
        onChange={(e) => handleImposterChange("three", e.target.value)}
        className="flex-2 range range-primary"
        min="0"
        max="100"
      />
      <span className="ml-4 w-12 text-center">{imposterChances.three}%</span>
    </div>
    <div>
      <label htmlFor="category" className="text-sm font-semibold">
        Category:
      </label>
      <input
        type="text"
        id="category"
        className="bg-gray-700 text-white p-2 rounded w-full"
        value={category}
        onChange={(e) => setCategory(e.target.value)}
      />
    </div>
  </div>
);
