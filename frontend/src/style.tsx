import { createGlobalStyle } from 'styled-components'

const GlobalStyles = createGlobalStyle`
  :root {
    --primary-color: #3498db;
    --secondary-color: #2ecc71;
    --font-size: 15px;
    --font-color: rgb(230, 230, 230);
    --background-color: #0D0D0D;
    --nav-background-color: #141414;
  }

  /* Reset CSS */
  *, *::before, *::after {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    background-color: var(--background-color); /* Cambia el uso aquí */
    font-size: var(--font-size);
    color: var(--font-color);

    font-family: "Inter", sans-serif;
    font-optical-sizing: auto;
    font-weight: 200;
    font-style: normal;
  }

  /* Link styles */
  a {
    text-decoration: none;
    color: inherit;
    transition: color 0.3s ease;

    &:hover {
      color: #007BFF;
    }
  }

  button {
    border: none;
    outline: none;
    cursor: pointer;
    font-family: inherit;
    color: var(--font-color);
    background-color: #333333;
    padding: 10px 20px;
    border-radius: 10px;
    text-transform: capitalize;
  }


  .row {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }
`;

export default GlobalStyles;