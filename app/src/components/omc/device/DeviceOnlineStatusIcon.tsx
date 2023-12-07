import { CheckCircleTwoTone, CloseCircleTwoTone } from "@ant-design/icons";
import { Tooltip } from "antd";

interface DeviceStatusIconProps {
  device?: API.Device;
}

const DeviceStatusIcon: React.FC<DeviceStatusIconProps> = ({
  device,
}) => {
  if (device?.Online) {
    return (
      <Tooltip title={'Online'}>
        <CheckCircleTwoTone twoToneColor="#52c41a" />
      </Tooltip>
    )
  }
  return (
    <Tooltip title={'Offline'}>
      <CloseCircleTwoTone twoToneColor="#eb2f96" />
    </Tooltip>
  )
};

export default DeviceStatusIcon;
