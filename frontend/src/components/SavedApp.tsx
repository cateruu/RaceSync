import React, { useEffect, useState } from 'react';
import { LoadImage } from '../../wailsjs/go/fileService/FileService';
import TrashIcon from '../assets/Trash.png';

export interface AppData {
  name: string;
  icon: string;
  path: string;
}

interface Props {
  data: AppData;
  onRemove: (name: string) => void;
}

const SavedApp = ({ data, onRemove }: Props) => {
  const [imageSrc, setImageSrc] = useState<string>('');

  useEffect(() => {
    (async () => {
      try {
        const imageSrc = await LoadImage(data.icon);
        setImageSrc(imageSrc);
      } catch (error) {
        console.error(error);
      }
    })();
  }, [data.icon]);

  return (
    <div className='h-10 w-full flex group'>
      <div className='min-w-10 w-10 h-10 p-1 bg-blue-950 rounded-lg flex items-center justify-center mr-3'>
        <img src={imageSrc} alt='app logo' className='w-full ' />
      </div>
      <div className='flex flex-col justify-between'>
        <p className='text-sm'>{`${data.name[0].toUpperCase()}${data.name.slice(
          1
        )}`}</p>
        <p
          className='text-slate-600 text-[10px] w-[270px] overflow-hidden whitespace-nowrap text-nowrap text-ellipsis'
          title={data.path.replaceAll('\\', '/')}
        >
          {data.path.replaceAll('\\', '/')}
        </p>
      </div>
      <div
        onClick={() => onRemove(data.name)}
        className='ml-auto self-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer w-5 h-5 flex items-center justify-center'
      >
        <img src={TrashIcon} alt='Delete' />
      </div>
    </div>
  );
};

export default SavedApp;
