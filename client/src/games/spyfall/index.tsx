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
import { LocationModal } from "./location-modal";
import { GameControl } from "./game-controls";
import { QuestionsModal } from "./questions-modal";

export type IPlayer = {
  id: string;
  name: string;
  points: number;
};

export function Spyfall() {
  const { user, isAdmin, isLoading } = useAuth();
  const router = useRouter();

  // Game state
  const [loadingGameStatus, setLoadingGameStatus] = useState(true);
  const [isPlaying, setIsPlaying] = useState(false);
  const [players, setPlayers] = useState([]);
  const [location, setLocation] = useState("");
  const [role, setRole] = useState("");
  const [gameStarted, setGameStarted] = useState(false);

  // Game setup
  const [spiesChances, setSpiesChances] = useState({
    one: 100,
    two: 0,
    three: 0,
  });

  // WebSocket connection
  const [ws, setWs] = useState<WebSocket | null>(null);

  const getGameStatus = async () => {
    if (!isLoading && user) {
      try {
        const gameStatus = await api.getSpyfallGameStatus(user.id)

        setGameStarted(gameStatus.gameStarted);
        setIsPlaying(gameStatus.inGame);
        setRole(gameStatus.role);
        setLocation(gameStatus.location);
        setLoadingGameStatus(false);
      } catch (error) {
        console.error("Error getting game status: ", error);
        setLoadingGameStatus(false);
      }
    }
  }

  useEffect(() => {
    const initializeWebSocket = () => {
      const newWs = new WebSocket(
        `${process.env.NEXT_PUBLIC_WS_API_URL}/games/spyfall/ws`
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
        getGameStatus()
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
      getGameStatus()
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
        setRole(data.role);
        setLocation(data.location);
        setGameStarted(true);
        setIsPlaying(true);
        break;
      case "resetGame":
        setGameStarted(false);
        setRole("");
        setLocation("");
        break;
      case "removedPlayer":
        router.push("/games", { scroll: false });
        break;
      case "winner":
        const spyWon = data.spyWon as boolean;
        if (spyWon) {
          toast.success("Spy venceu!", {
            duration: 3000,
          });
        } else {
          toast.success("Time venceu!", {
            duration: 3000,
          });
        }
        break;
      case "spiesNumber":
        toast.success(`Há ${data.spiesNumber} espiões na partida.`, {
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
          ...spiesChances,
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

  const handleRoundWinner = (spyWon: boolean) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "decideWinner", spyWon }));
    }
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

  const showSpiesNumber = () => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "showSpiesNumber" }));
    }
  };

  const handleNumberOfSpiesChange = (key: string, value: string) => {
    const newChances = { ...spiesChances, [key]: parseInt(value) };
    setSpiesChances(newChances);
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
            <>
              <GameStarted role={role} location={location} inGame={isPlaying} />
              {isAdmin && (
                <div className="bg-gray-900 text-white p-8 rounded-lg shadow-lg max-w-sm">
                  <GameControl
                    resetGame={resetGame}
                    decideWinner={handleRoundWinner}
                    showSpiesNumber={showSpiesNumber}
                  />
                </div>
              )}
            </>
          ) : (
            <GameHome
              players={players}
              startGame={startGame}
              isAdmin={isAdmin}
              removePlayer={removePlayer}
            />
          )}
          <LocationModal />
          <QuestionsModal />
          {isAdmin && !gameStarted && (
            <GameSetup
              spiesChances={spiesChances}
              handleSpiesChange={handleNumberOfSpiesChange}
              resetPoints={resetPoints}
            />
          )}
        </div>
      </div>
    </main>
  );
}
