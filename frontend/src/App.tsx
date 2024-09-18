import PlusCirle from './assets/PlusCircle.png';
import { OpenFile, GetAppsData } from '../wailsjs/go/fileService/FileService';
import { useEffect, useState } from 'react';
import Spinner from './loaders/Spinner';
import SavedApp, { AppData } from './components/SavedApp';

interface Data {
  [key: string]: AppData;
}

function App() {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const [data, setData] = useState<Data | null>(null);

  const openFile = async () => {
    setIsLoading(true);

    try {
      const data = await OpenFile();
      setData(data);
    } catch (error) {
      setError(error as string);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    const getData = async () => {
      try {
        const data = await GetAppsData();
        setData(data);
      } catch (error) {
        setError(error as string);
      }
    };

    getData();
  }, []);

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
          {error && <p>{error}</p>}
          <h2 className='mt-4 mb-5'>Added apps</h2>
          {data &&
            Object.entries(data).map(([_, data]) => <SavedApp data={data} />)}
        </div>
      </div>
    </main>
  );
}

export default App;
