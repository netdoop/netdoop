import { useEffect, useState } from 'react';
import services from '@/services/netdoop';
import dayjs from 'dayjs';

export const useSystemTime = () => {
  const [systemTimeDiff, setSystemTimeDiff,] = useState<number>(0); 

  const fetchSystemTime = async () => {
    try {
      const response = await services.system.getSystemTime();
      if (response?.current) {
        const diff = dayjs(Date.now()).unix() - response.current
        setSystemTimeDiff(diff);
      }
    } catch (error) {
      // Handle error here
      console.error('Error fetching system time:', error);
    }
  };

  useEffect(() => {
    // Fetch initial system time
    fetchSystemTime();
  }, []);

  return {systemTimeDiff};
};