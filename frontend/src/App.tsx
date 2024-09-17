import PlusCirle from './assets/PlusCircle.png';

function App() {
  return (
    <main className='h-screen w-screen text-white flex flex-col items-center gap-4 font-azeretMono'>
      <h1 className='font-fasterOne text-6xl'>RaceSync</h1>
      <div className='w-[340px]'>
        <div className='flex flex-col gap-1 w-full'>
          <button className='w-full h-11 text-sm font-medium rounded-xl bg-emerald-950 border border-green-700'>
            Launch apps
          </button>
          <button className='w-full h-11 text-sm font-medium rounded-xl bg-blue-950 border border-blue-800 flex justify-center items-center gap-[10px]'>
            <img src={PlusCirle} alt='plus circle' /> New app
          </button>
        </div>
      </div>
    </main>
  );
}

export default App;
