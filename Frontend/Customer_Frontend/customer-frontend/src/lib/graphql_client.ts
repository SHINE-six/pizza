import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';
import { getCookie } from 'cookies-next';
// import { useUserStore } from '../../store/store';

const createApolloClient = () => {
  // const { username, email } = useUserStore() as { username: string, email: string };
  const url = process.env.API_GATEWAY || 'https://xqtvgb1k-8080.asse.devtunnels.ms';
  const uri = `${url}/graphql`;
  // Create an HTTP link to the GraphQL server

  const httpLink = new HttpLink({
    uri: uri,
  });

  // Use setContext to create a middleware for the Apollo Link chain.
  // This allows you to dynamically set headers.
  const authLink = setContext((_, { headers }) => {
		const token = getCookie('token');
    console.log(token);
    // Return the headers to the context so httpLink can read them
    return {
      headers: {
        ...headers,
        // Add your headers here
        'Authorization': token,
        // Add any other headers you need
      }
    };
  });

  // Combine the authLink and httpLink. The order is important;
  // authLink must come before httpLink because Apollo Link
  // executes middleware in order.
  return new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });
};

export const apolloClient = createApolloClient();
