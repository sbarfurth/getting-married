import { createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { GuestService } from '../../gen/guest/v1/guest_service_pb';

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_URL,
});
export const client = createClient(GuestService, transport);
