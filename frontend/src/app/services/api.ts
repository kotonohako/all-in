import axios from 'axios';

// GETリクエスト
export const getJSON = async (): Promise<any> => {
  const url = "http://localhost:8080/v1/kotonohas";
  const response = await axios.get(url);
  return response.data;
};