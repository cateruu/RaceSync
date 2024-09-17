import PlusCirle from './assets/PlusCircle.png';
import { OpenFile } from '../wailsjs/go/fileService/FileService';
import { useState } from 'react';
import Spinner from './loaders/Spinner';

function App() {
  const [file, setFile] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const openFile = async () => {
    setIsLoading(true);
    const file = await OpenFile();
    setFile(file);

    setIsLoading(false);
  };

  return (
    <main className='h-screen w-screen text-white flex flex-col items-center gap-4 font-azeretMono'>
      <h1 className='font-fasterOne text-6xl'>RaceSync</h1>
      <div className='w-[340px]'>
        <div className='flex flex-col gap-1 w-full'>
          <button className='w-full h-11 text-sm font-medium rounded-xl bg-emerald-950 border border-green-700'>
            Launch apps
          </button>
          <button
            className='w-full h-11 text-sm font-medium rounded-xl bg-blue-950 border border-blue-800 flex justify-center items-center gap-[10px]'
            onClick={openFile}
          >
            {isLoading ? (
              <Spinner />
            ) : (
              <>
                <img src={PlusCirle} alt='plus circle' /> New app
              </>
            )}
          </button>
          {file}
        </div>
      </div>
    </main>
  );
}

export default App;
