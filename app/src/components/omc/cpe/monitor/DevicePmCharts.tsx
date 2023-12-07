import { ProCard } from "@ant-design/pro-components";
import { Tabs } from "antd";
import dayjs, { Dayjs } from "dayjs";
import { useState } from "react";
import { DevicePmValueChart } from "../../device";
import { DateRangePick } from "../../common";

interface Props {
  device: API.Device;
};
const DevicePmCharts: React.FC<Props> = ({ device }) => {
  const [timeRange, setTimeRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([dayjs(), dayjs()]);
  const handleTimeRangeChange = (value: [Dayjs, Dayjs]) => {
    setTimeRange(value);
  }
  const chartTabItems = [
    {
      key: "UL_MCS",
      label: "UL_MCS",
      children: <DevicePmValueChart names={["ULMCS", "NR_ULMCS"]} productType="cpe" device={device}
        dateRange={timeRange} height={320} />
    },
    {
      key: "DL_MCS",
      label: "DL_MCS",
      children: <DevicePmValueChart names={["DLMCS", "NR_DLMCS"]} productType="cpe" device={device}
        dateRange={timeRange} height={320} />
    },
    {
      key: "PCI",
      label: "PCI",
      children: <DevicePmValueChart names={["PCI", "NR_PCI"]} productType="cpe" device={device}
        dateRange={timeRange} height={320} />
    },
    {
      key: "RSRP",
      label: "RSRP",
      children: <DevicePmValueChart names={["RSRP", "RSRP0", "RSRP1", "SSB_RSRP", "SSB_RSRP0", "SSB_RSRP1"]} productType="cpe" device={device}
        unit="dBm" dateRange={timeRange} height={320} />
    },
    {
      key: "RSRQ",
      label: "RSRQ",
      children: <DevicePmValueChart names={["RSRQ", "SSB_RSRQ"]} productType="cpe" device={device}
        unit="dBm" dateRange={timeRange} height={320} />
    },
    {
      key: "SINR",
      label: "SINR",
      children: <DevicePmValueChart names={["SINR", "SSB_SINR"]} productType="cpe" device={device}
        unit="dB" dateRange={timeRange} height={320} />
    },

    {
      key: "DL Throughput",
      label: "DL Throughput",
      children: <DevicePmValueChart names={["TotalDownload"]} productType="cpe" device={device}
        unit="Bytes" dateRange={timeRange} height={320} func='increase' />
    },
    {
      key: "UL Throughput",
      label: "UL Throughput",
      children: <DevicePmValueChart names={["TotalUpload"]} productType="cpe" device={device}
        unit="Bytes" dateRange={timeRange} height={320} func='increase' />
    },
  ];
  return (
    <ProCard layout="center" direction="column">
      <ProCard >
        <DateRangePick onChange={handleTimeRangeChange} />
      </ProCard>
      <ProCard  >
        <Tabs defaultActiveKey="Network" tabPosition="left" items={chartTabItems} />)
      </ProCard>
    </ProCard>
  )
};

export default DevicePmCharts;