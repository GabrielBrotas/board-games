import axios from "axios";

const baseAPI = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});

type CreateUserResponse = {
  id: string;
  name: string;
};

const createUserOrLogin = async (name: string): Promise<CreateUserResponse> => {
  const response = await baseAPI.post(`/login`, {
    name,
  });

  return response.data;
};

type GameStatusResponse = {
  gameStarted: boolean;
  word: string;
  inGame: boolean;
};
const getImposterGameStatus = async (
  userID: string
): Promise<GameStatusResponse> => {
  const response = await baseAPI.get(`/game-status?u=${userID}`);
  return response.data;
};

export const api = {
  createUserOrLogin,
  getImposterGameStatus,
};
