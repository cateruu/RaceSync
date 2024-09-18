import PlusCirle from './assets/PlusCircle.png';
import {
  OpenFile,
  GetAppsData,
  RemoveApp,
  LaunchApps,
} from '../wailsjs/go/fileService/FileService';
import { useEffect, useState } from 'react';
import Spinner from './loaders/Spinner';
import SavedApp, { AppData } from './components/SavedApp';

interface Data {
  [key: string]: AppData;
}

function App() {
  const [isOpeningLoading, setIsOpeningLoading] = useState(false);
  const [isLaunchingLoading, setIsLaunchingLoading] = useState(false);
  const [error, setError] = useState('');
  const [data, setData] = useState<Data | null>(null);

  const openFile = async () => {
    setIsOpeningLoading(true);

    try {
      const data = await OpenFile();
      setData(data);
    } catch (error) {
      setError(error as string);
    } finally {
      setIsOpeningLoading(false);
    }
  };

  const removeApp = async (name: string) => {
    try {
      const data = await RemoveApp(name);
      setData(data);
    } catch (error) {
      setError(error as string);
    }
  };

  const launchApps = async () => {
    setIsLaunchingLoading(true);

    try {
      await LaunchApps();
    } catch (error) {
      setError(error as string);
    } finally {
      setIsLaunchingLoading(false);
    }
  };

  useEffect(() => {
    const getInitData = async () => {
      try {
        const data = await GetAppsData();
        setData(data);
      } catch (error) {
        if (typeof error === 'string') {
          if (error !== 'unable to read data file') {
            setError(error);
          }
        }
      }
    };

    getInitData();
  }, []);

  return (
    <main className='h-screen w-screen text-white flex flex-col items-center gap-4 font-azeretMono'>
      <h1 className='font-fasterOne text-6xl'>RaceSync</h1>
      <div className='w-[340px]'>
        <div className='flex flex-col gap-1 w-full'>
          <button
            className='w-full h-11 text-sm font-medium rounded-xl bg-emerald-950 border border-green-700'
            onClick={launchApps}
          >
            {isLaunchingLoading ? (
              <Spinner color='fill-green-500' />
            ) : (
              <>Launch apps</>
            )}
          </button>
          <button
            className='w-full h-11 text-sm font-medium rounded-xl bg-blue-950 border border-blue-800 flex justify-center items-center gap-[10px]'
            onClick={openFile}
          >
            {isOpeningLoading ? (
              <Spinner color='fill-blue-500' />
            ) : (
              <>
                <img src={PlusCirle} alt='plus circle' /> New app
              </>
            )}
          </button>
          {error && <p>{error}</p>}
          {data && Object.entries(data).length > 0 && (
            <h2 className='mt-2 mb-3'>Added apps</h2>
          )}
          <section className='flex flex-col gap-2'>
            {data &&
              Object.entries(data).map(([_, data]) => (
                <SavedApp data={data} onRemove={removeApp} />
              ))}
          </section>
        </div>
      </div>
    </main>
  );
}

export default App;
