"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import { useAuth } from "@/context/AuthContext";
import { GameHome } from "./game-home";
import { GameStarted } from "./game-started";
import { GameSetup } from "./admin-game-setup";
import { useRouter } from "next/navigation";
import { Loading } from "@/components/loading";
import { toast } from "react-hot-toast";
import { api } from "@/lib/api";

export type IPlayer = {
  id: string;
  name: string;
  points: number;
};

export function WhoIsTheImposter() {
  const { user, isAdmin, isLoading } = useAuth();
  const router = useRouter();

  // Game state
  const [loadingGameStatus, setLoadingGameStatus] = useState(true);
  const [isPlaying, setIsPlaying] = useState(false);
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
    const initializeWebSocket = () => {
      const newWs = new WebSocket(
        `${process.env.NEXT_PUBLIC_WS_API_URL}/games/impostor/ws`
      );
      setWs(newWs);

      newWs.onopen = () => handleWebSocketOpen(newWs);
      newWs.onmessage = handleWebSocketMessage;
      newWs.onclose = () => console.log("Connection closed");

      return newWs;
    };

    const newWs = initializeWebSocket();

    const handleVisibilityChange = () => {
      if (document.visibilityState === "visible" && (!ws || ws.readyState !== WebSocket.OPEN)) {
        console.log("Reconnecting WebSocket...");
        setWs(initializeWebSocket());
      }
    };

    document.addEventListener("visibilitychange", handleVisibilityChange);

    return () => {
      console.log("Closing connection...");
      newWs.close();
      document.removeEventListener("visibilitychange", handleVisibilityChange);
    };
  }, [user]);

  useEffect(() => {
    if (!isLoading && user) {
      api
        .getImposterGameStatus(user.id)
        .then((data) => {
          setGameStarted(data.gameStarted);
          setMessage(data.word);
          setIsPlaying(data.inGame);
          setLoadingGameStatus(false);
        })
        .catch((error) => {
          console.error("Error getting game status: ", error);
          setLoadingGameStatus(false);
        });
    }
  }, [user, isLoading]);

  const handleWebSocketOpen = (newWs: WebSocket) => {
    if (newWs.readyState === WebSocket.OPEN && user) {
      newWs.send(JSON.stringify({ type: "connected", id: user.id }));
    } else {
      console.error("WebSocket is not open.");
    }
  };

  const handleWebSocketMessage = (event: MessageEvent<any>) => {
    const data = JSON.parse(event.data);
    switch (data.type) {
      case "playerList":
        setPlayers(data.players);
        break;
      case "role":
        setMessage(data.wordOrRole);
        setGameStarted(true);
        setIsPlaying(true);
        break;
      case "resetGame":
        setGameStarted(false);
        setMessage("");
        break;
      case "removedPlayer":
        router.push("/games", { scroll: false });
        break;
      case "winner":
        const impostorsWon = data.impostorsWon as boolean;
        if (impostorsWon) {
          toast.success("Impostores venceram!", {
            duration: 3000,
          });
        } else {
          toast.success("Tripulantes venceram!", {
            duration: 3000,
          });
        }
        break;
      case "impostorsNumber":
        toast.success(`Há ${data.impostorsNumber} impostores na partida.`, {
          duration: 3000,
        });
        break;
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

  const handleNumberOfImposterChange = (key: string, value: string) => {
    const newChances = { ...imposterChances, [key]: parseInt(value) };
    setImposterChances(newChances);
  };

  const removePlayer = (id: string) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      console.log("Removing player with id: ", id);
      ws.send(JSON.stringify({ type: "removePlayer", id: id }));
    }
  };

  const resetPoints = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "resetPoints" }));
    }
  };

  const showImpostorsNumber = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "showImpostorsNumber" }));
    }
  };

  if (isLoading || loadingGameStatus) {
    return <Loading />;
  }

  return (
    <main className="bg-gray-800 min-h-screen p-4">
      <div>
        <Link href="/games">
          <b className="absolute top-0 left-0 m-4 bg-gray-500 text-white px-4 py-2 rounded">
            Voltar
          </b>
        </Link>
      </div>

      <div className="min-h-screen flex flex-col items-center justify-center gap-4">
        <div className="flex flex-col gap-4 w-full max-w-sm px-4">
          {gameStarted ? (
            <GameStarted
              message={message}
              isAdmin={isAdmin}
              resetGame={resetGame}
              decideWinner={handleRoundWinner}
              inGame={isPlaying}
              showImpostorsNumber={showImpostorsNumber}
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
              handleImposterChange={handleNumberOfImposterChange}
              category={category}
              setCategory={setCategory}
              difficulty={difficulty}
              setDifficulty={setDifficulty}
              resetPoints={resetPoints}
            />
          )}
        </div>
      </div>
    </main>
  );
}
