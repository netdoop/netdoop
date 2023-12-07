import { CheckCircleTwoTone, CloseCircleTwoTone } from "@ant-design/icons";
import { Tooltip } from "antd";

interface DeviceStatusIconProps {
  device?: API.Device;
}

const DeviceStatusIcon: React.FC<DeviceStatusIconProps> = ({
  device,
}) => {
  if (device?.ActiveStatus === 'active') {
    return (
      <Tooltip title={'Active'}>
        <CheckCircleTwoTone twoToneColor="#52c41a" />
        <span>Active</span>
      </Tooltip>
    )
  }
  return (
    <Tooltip title={'Inactive'}>
      <CloseCircleTwoTone twoToneColor="#eb2f96" />
      <span>Inactive</span>
      </Tooltip>
  )
};

export default DeviceStatusIcon;
