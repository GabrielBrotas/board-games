"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import { useAuth } from "@/context/AuthContext";
import { GameHome } from "./game-home";
import { GameStarted } from "./game-started";
import { GameSetup } from "./admin-game-setup";
import { useRouter } from "next/navigation";

export type IPlayer = {
  id: string;
  name: string;
  points: number;
};

export function WhoIsTheImposter() {
  const { user, isAdmin, isLoading } = useAuth();
  const router = useRouter();

  // Game state
  const [players, setPlayers] = useState([]);
  const [message, setMessage] = useState("");
  const [gameStarted, setGameStarted] = useState(false);

  // Game setup
  const [imposterChances, setImposterChances] = useState({
    one: 100,
    two: 0,
    three: 0,
  });
  const [category, setCategory] = useState("");
  const [difficulty, setDifficulty] = useState("");

  // WebSocket connection
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const newWs = new WebSocket("ws://localhost:8081/ws");
    setWs(newWs);

    newWs.onopen = () => handleWebSocketOpen(newWs);
    newWs.onmessage = handleWebSocketMessage;
    newWs.onclose = () => console.log("Connection closed");

    return () => {
      console.log("Closing connection...");
      newWs.close();
    };
  }, []);

  useEffect(() => {
    if (!isLoading && ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "changeName", newName: user }));
    }
  }, [user, isLoading, ws?.readyState]);

  const handleWebSocketOpen = (newWs: WebSocket) => {
    console.log("Connected to server");
    if (newWs.readyState === WebSocket.OPEN) {
      newWs.send(JSON.stringify({ type: "changeName", newName: user }));
    } else {
      console.error("WebSocket is not open.");
    }
  };

  const handleWebSocketMessage = (event: MessageEvent<any>) => {
    const data = JSON.parse(event.data);
    console.log(`Received message: ${data.type}`);
    switch (data.type) {
      case "playerList":
        setPlayers(data.players);
        break;
      case "role":
        setMessage(data.wordOrRole);
        setGameStarted(true);
        break;
      case "resetGame":
        setGameStarted(false);
        setMessage("");
        break;
      case "removedPlayer":
        router.push("/", { scroll: false });
      default:
        break;
    }
  };

  const startGame = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(
        JSON.stringify({
          type: "startGame",
          category,
          difficulty,
          ...imposterChances,
        })
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

  const handleRoundWinner = (impostorWon: boolean) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(
        JSON.stringify({ type: "decideWinner", impostorWon: impostorWon })
      );
    }
  };

  const handleImposterChange = (key: string, value: string) => {
    const newChances = { ...imposterChances, [key]: parseInt(value) };
    setImposterChances(newChances);
  };

  const removePlayer = (id: string) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      console.log("Removing player with id: ", id);
      ws.send(JSON.stringify({ type: "removePlayer", playerID: id }));
    }
  };

  const resetPoints = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "resetPoints" }));
    }
  };

  return (
    <div className="bg-gray-800 min-h-screen flex flex-col items-center justify-center gap-4">
      <Link href="/">
        <b className="absolute top-0 left-0 m-4 bg-gray-500 text-white px-4 py-2 rounded">
          Voltar
        </b>
      </Link>
      <div className="flex flex-col gap-4 w-full max-w-sm px-4">
        {gameStarted ? (
          <GameStarted
            message={message}
            isAdmin={isAdmin}
            resetGame={resetGame}
            decideWinner={handleRoundWinner}
          />
        ) : (
          <GameHome
            players={players}
            startGame={startGame}
            isAdmin={isAdmin}
            removePlayer={removePlayer}
          />
        )}
        {isAdmin && !gameStarted && (
          <GameSetup
            imposterChances={imposterChances}
            handleImposterChange={handleImposterChange}
            category={category}
            setCategory={setCategory}
            difficulty={difficulty}
            setDifficulty={setDifficulty}
            resetPoints={resetPoints}
          />
        )}
      </div>
    </div>
  );
}
