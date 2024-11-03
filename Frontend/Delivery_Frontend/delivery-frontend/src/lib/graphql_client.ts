import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';
// import { useUserStore } from '../../store/store';

const createApolloClient = () => {
  // const { username, email } = useUserStore() as { username: string, email: string };
  const url = process.env.API_GATEWAY || 'https://xqtvgb1k-8080.asse.devtunnels.ms';
  const uri = `${url}/graphql`;
  // Create an HTTP link to the GraphQL server

  const httpLink = new HttpLink({
    uri: uri,
  });


  // Combine the authLink and httpLink. The order is important;
  // authLink must come before httpLink because Apollo Link
  // executes middleware in order.
  return new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache(),
  });
};

export const apolloClient = createApolloClient();
