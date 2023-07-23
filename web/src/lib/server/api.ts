import { env } from '$env/dynamic/private';
import { AuthAPIClient } from '$gen/auth/auth.client';
import { StationAPIClient } from '$gen/station/station.client';
import { ChannelCredentials } from '@grpc/grpc-js';
import { GrpcTransport } from '@protobuf-ts/grpc-transport';

const transport = new GrpcTransport({
	host: env.TRAIN_API_BASE_URL ?? 'localhost:3000',
	channelCredentials: ChannelCredentials.createInsecure()
});

export const stationClient = new StationAPIClient(transport);

export const authClient = new AuthAPIClient(transport);
