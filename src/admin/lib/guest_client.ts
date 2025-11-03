import { createClient, type Interceptor } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { GuestService } from '../../../gen/guest/v1/guest_service_pb';
import { getToken } from './auth';

const ensureAuth: Interceptor = (next) => async (req) => {
  const token = getToken();
  if (token) {
    req.header.set('Authorization', token);
  }
  return await next(req);
};

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_URL,
  interceptors: [ensureAuth],
});
export const client = createClient(GuestService, transport);
