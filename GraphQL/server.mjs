import { ApolloServer } from 'apollo-server';
import { ApolloGateway, IntrospectAndCompose, RemoteGraphQLDataSource } from "@apollo/gateway";

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
    willSendRequest({ request, context }) {
        request.http.headers.set('Authorization', context.authToken);
        if (context.authToken) {
            console.log('Setting Authorization header with token:', context.authToken);
        }
    }
  }

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: 'order', url: 'http://order_service-backend:10072/query' },
            { name: 'menu', url: 'http://menu_service-backend:10071/query' },
    ],
    }),
    buildService({ name, url }) {
        return new AuthenticatedDataSource({ url });
    }
});

const server = new ApolloServer({ 
    gateway, 
    subscriptions: false,
    context: ({ req }) => {
        const authToken = req.headers.authorization || '';
        return { authToken };
    },
    context: ({ req }) => ({ authToken: req.headers.authorization || '' })
});

server.listen({ port: 10070 }).then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});