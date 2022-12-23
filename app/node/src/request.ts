import Axios, {AxiosResponse, AxiosRequestConfig} from 'axios';

export const request = (url: string): Promise<AxiosResponse> => {
  const config: AxiosRequestConfig = {
    method: 'GET',
    url: `http://localhost:26657/${url}`,
  };

  return Axios.request(config);
};
