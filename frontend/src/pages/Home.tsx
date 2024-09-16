import { Link } from 'react-router-dom';

function App() {
  const openOverlay = async () => {};

  return (
    <div className='h-screen w-screen text-white' onClick={() => {}}>
      <h1 className='text-3xl text-red-700 pl-3'>iRacing Utlity</h1>
      <button onClick={openOverlay}>Open Overlay</button>
      <br></br>
      <Link to='/overlay'>Overlay</Link>
    </div>
  );
}

export default App;
