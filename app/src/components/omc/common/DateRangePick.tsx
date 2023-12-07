import React, { useEffect } from 'react';
import { DatePicker } from 'antd';
import { timestampToDayjs } from '@/utils/format';
import { RangePickerProps } from "antd/es/date-picker";

import dayjs from 'dayjs';

const { RangePicker } = DatePicker;

interface Props {
  value?: [dayjs.Dayjs, dayjs.Dayjs],
  onChange?: (value: [dayjs.Dayjs, dayjs.Dayjs],) => void;
}

const DateRangePick: React.FC<Props> = ({
  value,
  onChange,
}) => {
  const defaultFrom = dayjs().startOf('h').add(-23, 'h');
  const defaultTo = dayjs().startOf('h').add(0, 'h');
  // const [timeRange, setTimeRange] = useState<[dayjs.Dayjs,dayjs.Dayjs]>();
  useEffect(() => {
    if (value === undefined) {
      if (onChange) {
        const from = defaultFrom.unix();
        const to = defaultTo.unix();
        onChange([timestampToDayjs(from), timestampToDayjs(to)])
      }
    }
  }, [value]);

  const onOk = (value: RangePickerProps['value']) => {
    if (value && value[0] && value[1]) {
      const from = value[0]?.unix();
      const to = value[1]?.unix();
      if (onChange) {
        onChange([timestampToDayjs(from), timestampToDayjs(to)])
      }
    }
  };
  const rangePresets: {
    label: string;
    value: [dayjs.Dayjs, dayjs.Dayjs];
  }[] = [
      { label: 'Last 6 Hours', value: [dayjs().startOf('h').add(-5, 'h'), dayjs().startOf('h').add(0, 'h')] },
      { label: 'Last 12 Hours', value: [dayjs().startOf('h').add(-11, 'h'), dayjs().startOf('h').add(0, 'h')] },
      { label: 'Last 24 Hours', value: [dayjs().startOf('h').add(-23, 'h'), dayjs().startOf('h').add(0, 'h')] },
      { label: 'Last 2 Days', value: [dayjs().startOf('d').add(-1, 'd'), dayjs().startOf('d').add(0, 'd')] },
      { label: 'Last 7 Days', value: [dayjs().startOf('d').add(-6, 'd'), dayjs().startOf('d').add(0, 'd')] },
      { label: 'Last 30 Days', value: [dayjs().startOf('d').add(-29, 'd'), dayjs().startOf('d').add(0, 'd')] },
    ];

  return (
    <RangePicker
      defaultValue={value||[defaultFrom, defaultTo]}
      presets={rangePresets}
      showTime={{ format: 'HH:mm' }}
      format="YYYY-MM-DD HH:mm"
      onChange={onOk}
      onOk={onOk}
    />
  );
};

export default DateRangePick;