#ifndef _TINYROS_ROSTOPIC_SUBSCRIBERS_h
#define _TINYROS_ROSTOPIC_SUBSCRIBERS_h

#include "tiny_ros/ros.h"
#include "tiny_ros/actionlib_msgs/GoalID.h"
#include "tiny_ros/actionlib_msgs/GoalStatus.h"
#include "tiny_ros/actionlib_msgs/GoalStatusArray.h"
#include "tiny_ros/diagnostic_msgs/DiagnosticArray.h"
#include "tiny_ros/diagnostic_msgs/DiagnosticStatus.h"
#include "tiny_ros/diagnostic_msgs/KeyValue.h"
#include "tiny_ros/gazebo_msgs/ContactsState.h"
#include "tiny_ros/gazebo_msgs/ContactState.h"
#include "tiny_ros/gazebo_msgs/LinkState.h"
#include "tiny_ros/gazebo_msgs/LinkStates.h"
#include "tiny_ros/gazebo_msgs/ModelState.h"
#include "tiny_ros/gazebo_msgs/ModelStates.h"
#include "tiny_ros/gazebo_msgs/ODEJointProperties.h"
#include "tiny_ros/gazebo_msgs/ODEPhysics.h"
#include "tiny_ros/gazebo_msgs/WorldState.h"
#include "tiny_ros/geometry_msgs/Accel.h"
#include "tiny_ros/geometry_msgs/AccelStamped.h"
#include "tiny_ros/geometry_msgs/AccelWithCovariance.h"
#include "tiny_ros/geometry_msgs/AccelWithCovarianceStamped.h"
#include "tiny_ros/geometry_msgs/Inertia.h"
#include "tiny_ros/geometry_msgs/InertiaStamped.h"
#include "tiny_ros/geometry_msgs/Point.h"
#include "tiny_ros/geometry_msgs/Point32.h"
#include "tiny_ros/geometry_msgs/PointStamped.h"
#include "tiny_ros/geometry_msgs/Polygon.h"
#include "tiny_ros/geometry_msgs/PolygonStamped.h"
#include "tiny_ros/geometry_msgs/Pose.h"
#include "tiny_ros/geometry_msgs/Pose2D.h"
#include "tiny_ros/geometry_msgs/PoseArray.h"
#include "tiny_ros/geometry_msgs/PoseStamped.h"
#include "tiny_ros/geometry_msgs/PoseWithCovariance.h"
#include "tiny_ros/geometry_msgs/PoseWithCovarianceStamped.h"
#include "tiny_ros/geometry_msgs/Quaternion.h"
#include "tiny_ros/geometry_msgs/QuaternionStamped.h"
#include "tiny_ros/geometry_msgs/Transform.h"
#include "tiny_ros/geometry_msgs/TransformStamped.h"
#include "tiny_ros/geometry_msgs/Twist.h"
#include "tiny_ros/geometry_msgs/TwistStamped.h"
#include "tiny_ros/geometry_msgs/TwistWithCovariance.h"
#include "tiny_ros/geometry_msgs/TwistWithCovarianceStamped.h"
#include "tiny_ros/geometry_msgs/Vector3.h"
#include "tiny_ros/geometry_msgs/Vector3Stamped.h"
#include "tiny_ros/geometry_msgs/Wrench.h"
#include "tiny_ros/geometry_msgs/WrenchStamped.h"
#include "tiny_ros/map_msgs/OccupancyGridUpdate.h"
#include "tiny_ros/map_msgs/PointCloud2Update.h"
#include "tiny_ros/map_msgs/ProjectedMap.h"
#include "tiny_ros/map_msgs/ProjectedMapInfo.h"
#include "tiny_ros/nav_msgs/GetMapAction.h"
#include "tiny_ros/nav_msgs/GetMapActionFeedback.h"
#include "tiny_ros/nav_msgs/GetMapActionGoal.h"
#include "tiny_ros/nav_msgs/GetMapActionResult.h"
#include "tiny_ros/nav_msgs/GetMapFeedback.h"
#include "tiny_ros/nav_msgs/GetMapGoal.h"
#include "tiny_ros/nav_msgs/GetMapResult.h"
#include "tiny_ros/nav_msgs/GridCells.h"
#include "tiny_ros/nav_msgs/MapMetaData.h"
#include "tiny_ros/nav_msgs/OccupancyGrid.h"
#include "tiny_ros/nav_msgs/Odometry.h"
#include "tiny_ros/nav_msgs/Path.h"
#include "tiny_ros/rosgraph_msgs/Clock.h"
#include "tiny_ros/rosgraph_msgs/Log.h"
#include "tiny_ros/rosgraph_msgs/TopicStatistics.h"
#include "tiny_ros/sensor_msgs/BatteryState.h"
#include "tiny_ros/sensor_msgs/CameraInfo.h"
#include "tiny_ros/sensor_msgs/ChannelFloat32.h"
#include "tiny_ros/sensor_msgs/CompressedImage.h"
#include "tiny_ros/sensor_msgs/FluidPressure.h"
#include "tiny_ros/sensor_msgs/Illuminance.h"
#include "tiny_ros/sensor_msgs/Image.h"
#include "tiny_ros/sensor_msgs/Imu.h"
#include "tiny_ros/sensor_msgs/JointState.h"
#include "tiny_ros/sensor_msgs/Joy.h"
#include "tiny_ros/sensor_msgs/JoyFeedback.h"
#include "tiny_ros/sensor_msgs/JoyFeedbackArray.h"
#include "tiny_ros/sensor_msgs/LaserEcho.h"
#include "tiny_ros/sensor_msgs/LaserScan.h"
#include "tiny_ros/sensor_msgs/MagneticField.h"
#include "tiny_ros/sensor_msgs/MultiDOFJointState.h"
#include "tiny_ros/sensor_msgs/MultiEchoLaserScan.h"
#include "tiny_ros/sensor_msgs/NavSatFix.h"
#include "tiny_ros/sensor_msgs/NavSatStatus.h"
#include "tiny_ros/sensor_msgs/PointCloud.h"
#include "tiny_ros/sensor_msgs/PointCloud2.h"
#include "tiny_ros/sensor_msgs/PointField.h"
#include "tiny_ros/sensor_msgs/Range.h"
#include "tiny_ros/sensor_msgs/RegionOfInterest.h"
#include "tiny_ros/sensor_msgs/RelativeHumidity.h"
#include "tiny_ros/sensor_msgs/Temperature.h"
#include "tiny_ros/sensor_msgs/TimeReference.h"
#include "tiny_ros/shape_msgs/Mesh.h"
#include "tiny_ros/shape_msgs/MeshTriangle.h"
#include "tiny_ros/shape_msgs/Plane.h"
#include "tiny_ros/shape_msgs/SolidPrimitive.h"
#include "tiny_ros/smach_msgs/SmachContainerInitialStatusCmd.h"
#include "tiny_ros/smach_msgs/SmachContainerStatus.h"
#include "tiny_ros/smach_msgs/SmachContainerStructure.h"
#include "tiny_ros/std_msgs/Bool.h"
#include "tiny_ros/std_msgs/Byte.h"
#include "tiny_ros/std_msgs/ByteMultiArray.h"
#include "tiny_ros/std_msgs/Char.h"
#include "tiny_ros/std_msgs/ColorRGBA.h"
#include "tiny_ros/std_msgs/Duration.h"
#include "tiny_ros/std_msgs/Empty.h"
#include "tiny_ros/std_msgs/Float32.h"
#include "tiny_ros/std_msgs/Float32MultiArray.h"
#include "tiny_ros/std_msgs/Float64.h"
#include "tiny_ros/std_msgs/Float64MultiArray.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/std_msgs/Int16.h"
#include "tiny_ros/std_msgs/Int16MultiArray.h"
#include "tiny_ros/std_msgs/Int32.h"
#include "tiny_ros/std_msgs/Int32MultiArray.h"
#include "tiny_ros/std_msgs/Int64.h"
#include "tiny_ros/std_msgs/Int64MultiArray.h"
#include "tiny_ros/std_msgs/Int8.h"
#include "tiny_ros/std_msgs/Int8MultiArray.h"
#include "tiny_ros/std_msgs/MultiArrayDimension.h"
#include "tiny_ros/std_msgs/MultiArrayLayout.h"
#include "tiny_ros/std_msgs/String.h"
#include "tiny_ros/std_msgs/Time.h"
#include "tiny_ros/std_msgs/UInt16.h"
#include "tiny_ros/std_msgs/UInt16MultiArray.h"
#include "tiny_ros/std_msgs/UInt32.h"
#include "tiny_ros/std_msgs/UInt32MultiArray.h"
#include "tiny_ros/std_msgs/UInt64.h"
#include "tiny_ros/std_msgs/UInt64MultiArray.h"
#include "tiny_ros/std_msgs/UInt8.h"
#include "tiny_ros/std_msgs/UInt8MultiArray.h"
#include "tiny_ros/stereo_msgs/DisparityImage.h"
#include "tiny_ros/tf/tfMessage.h"
#include "tiny_ros/tf2_msgs/LookupTransformAction.h"
#include "tiny_ros/tf2_msgs/LookupTransformActionFeedback.h"
#include "tiny_ros/tf2_msgs/LookupTransformActionGoal.h"
#include "tiny_ros/tf2_msgs/LookupTransformActionResult.h"
#include "tiny_ros/tf2_msgs/LookupTransformFeedback.h"
#include "tiny_ros/tf2_msgs/LookupTransformGoal.h"
#include "tiny_ros/tf2_msgs/LookupTransformResult.h"
#include "tiny_ros/tf2_msgs/TF2Error.h"
#include "tiny_ros/tf2_msgs/TFMessage.h"
#include "tiny_ros/tinyros_hello/TinyrosHello.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/trajectory_msgs/JointTrajectory.h"
#include "tiny_ros/trajectory_msgs/JointTrajectoryPoint.h"
#include "tiny_ros/trajectory_msgs/MultiDOFJointTrajectory.h"
#include "tiny_ros/trajectory_msgs/MultiDOFJointTrajectoryPoint.h"

namespace tinyros
{
template<typename MsgT>
class EchoSubscriber: public tinyros::Subscriber_
{
public:
MsgT msg;
  virtual void callback(unsigned char* data)
  {
    MsgT tmsg;
    tmsg.deserialize(data);
    spdlog::get("logger")->info("{0}[{1}]->>{2}", topic_, getMsgType(), tmsg.echo());
  }
  virtual std::string getMsgType()
  {
    return this->msg.getType();
  }
  virtual std::string getMsgMD5()
  {
    return this->msg.getMD5();
  }
  virtual int getEndpointType()
  {
    return tinyros_msgs::TopicInfo::ID_SUBSCRIBER;
  }
};

static std::map<std::string, tinyros::Subscriber_*> rostopic_subscribers = {
    {"actionlib_msgs/GoalID", new EchoSubscriber<actionlib_msgs::GoalID>()},
    {"actionlib_msgs/GoalStatus", new EchoSubscriber<actionlib_msgs::GoalStatus>()},
    {"actionlib_msgs/GoalStatusArray", new EchoSubscriber<actionlib_msgs::GoalStatusArray>()},
    {"diagnostic_msgs/DiagnosticArray", new EchoSubscriber<diagnostic_msgs::DiagnosticArray>()},
    {"diagnostic_msgs/DiagnosticStatus", new EchoSubscriber<diagnostic_msgs::DiagnosticStatus>()},
    {"diagnostic_msgs/KeyValue", new EchoSubscriber<diagnostic_msgs::KeyValue>()},
    {"gazebo_msgs/ContactsState", new EchoSubscriber<gazebo_msgs::ContactsState>()},
    {"gazebo_msgs/ContactState", new EchoSubscriber<gazebo_msgs::ContactState>()},
    {"gazebo_msgs/LinkState", new EchoSubscriber<gazebo_msgs::LinkState>()},
    {"gazebo_msgs/LinkStates", new EchoSubscriber<gazebo_msgs::LinkStates>()},
    {"gazebo_msgs/ModelState", new EchoSubscriber<gazebo_msgs::ModelState>()},
    {"gazebo_msgs/ModelStates", new EchoSubscriber<gazebo_msgs::ModelStates>()},
    {"gazebo_msgs/ODEJointProperties", new EchoSubscriber<gazebo_msgs::ODEJointProperties>()},
    {"gazebo_msgs/ODEPhysics", new EchoSubscriber<gazebo_msgs::ODEPhysics>()},
    {"gazebo_msgs/WorldState", new EchoSubscriber<gazebo_msgs::WorldState>()},
    {"geometry_msgs/Accel", new EchoSubscriber<geometry_msgs::Accel>()},
    {"geometry_msgs/AccelStamped", new EchoSubscriber<geometry_msgs::AccelStamped>()},
    {"geometry_msgs/AccelWithCovariance", new EchoSubscriber<geometry_msgs::AccelWithCovariance>()},
    {"geometry_msgs/AccelWithCovarianceStamped", new EchoSubscriber<geometry_msgs::AccelWithCovarianceStamped>()},
    {"geometry_msgs/Inertia", new EchoSubscriber<geometry_msgs::Inertia>()},
    {"geometry_msgs/InertiaStamped", new EchoSubscriber<geometry_msgs::InertiaStamped>()},
    {"geometry_msgs/Point", new EchoSubscriber<geometry_msgs::Point>()},
    {"geometry_msgs/Point32", new EchoSubscriber<geometry_msgs::Point32>()},
    {"geometry_msgs/PointStamped", new EchoSubscriber<geometry_msgs::PointStamped>()},
    {"geometry_msgs/Polygon", new EchoSubscriber<geometry_msgs::Polygon>()},
    {"geometry_msgs/PolygonStamped", new EchoSubscriber<geometry_msgs::PolygonStamped>()},
    {"geometry_msgs/Pose", new EchoSubscriber<geometry_msgs::Pose>()},
    {"geometry_msgs/Pose2D", new EchoSubscriber<geometry_msgs::Pose2D>()},
    {"geometry_msgs/PoseArray", new EchoSubscriber<geometry_msgs::PoseArray>()},
    {"geometry_msgs/PoseStamped", new EchoSubscriber<geometry_msgs::PoseStamped>()},
    {"geometry_msgs/PoseWithCovariance", new EchoSubscriber<geometry_msgs::PoseWithCovariance>()},
    {"geometry_msgs/PoseWithCovarianceStamped", new EchoSubscriber<geometry_msgs::PoseWithCovarianceStamped>()},
    {"geometry_msgs/Quaternion", new EchoSubscriber<geometry_msgs::Quaternion>()},
    {"geometry_msgs/QuaternionStamped", new EchoSubscriber<geometry_msgs::QuaternionStamped>()},
    {"geometry_msgs/Transform", new EchoSubscriber<geometry_msgs::Transform>()},
    {"geometry_msgs/TransformStamped", new EchoSubscriber<geometry_msgs::TransformStamped>()},
    {"geometry_msgs/Twist", new EchoSubscriber<geometry_msgs::Twist>()},
    {"geometry_msgs/TwistStamped", new EchoSubscriber<geometry_msgs::TwistStamped>()},
    {"geometry_msgs/TwistWithCovariance", new EchoSubscriber<geometry_msgs::TwistWithCovariance>()},
    {"geometry_msgs/TwistWithCovarianceStamped", new EchoSubscriber<geometry_msgs::TwistWithCovarianceStamped>()},
    {"geometry_msgs/Vector3", new EchoSubscriber<geometry_msgs::Vector3>()},
    {"geometry_msgs/Vector3Stamped", new EchoSubscriber<geometry_msgs::Vector3Stamped>()},
    {"geometry_msgs/Wrench", new EchoSubscriber<geometry_msgs::Wrench>()},
    {"geometry_msgs/WrenchStamped", new EchoSubscriber<geometry_msgs::WrenchStamped>()},
    {"map_msgs/OccupancyGridUpdate", new EchoSubscriber<map_msgs::OccupancyGridUpdate>()},
    {"map_msgs/PointCloud2Update", new EchoSubscriber<map_msgs::PointCloud2Update>()},
    {"map_msgs/ProjectedMap", new EchoSubscriber<map_msgs::ProjectedMap>()},
    {"map_msgs/ProjectedMapInfo", new EchoSubscriber<map_msgs::ProjectedMapInfo>()},
    {"nav_msgs/GetMapAction", new EchoSubscriber<nav_msgs::GetMapAction>()},
    {"nav_msgs/GetMapActionFeedback", new EchoSubscriber<nav_msgs::GetMapActionFeedback>()},
    {"nav_msgs/GetMapActionGoal", new EchoSubscriber<nav_msgs::GetMapActionGoal>()},
    {"nav_msgs/GetMapActionResult", new EchoSubscriber<nav_msgs::GetMapActionResult>()},
    {"nav_msgs/GetMapFeedback", new EchoSubscriber<nav_msgs::GetMapFeedback>()},
    {"nav_msgs/GetMapGoal", new EchoSubscriber<nav_msgs::GetMapGoal>()},
    {"nav_msgs/GetMapResult", new EchoSubscriber<nav_msgs::GetMapResult>()},
    {"nav_msgs/GridCells", new EchoSubscriber<nav_msgs::GridCells>()},
    {"nav_msgs/MapMetaData", new EchoSubscriber<nav_msgs::MapMetaData>()},
    {"nav_msgs/OccupancyGrid", new EchoSubscriber<nav_msgs::OccupancyGrid>()},
    {"nav_msgs/Odometry", new EchoSubscriber<nav_msgs::Odometry>()},
    {"nav_msgs/Path", new EchoSubscriber<nav_msgs::Path>()},
    {"rosgraph_msgs/Clock", new EchoSubscriber<rosgraph_msgs::Clock>()},
    {"rosgraph_msgs/Log", new EchoSubscriber<rosgraph_msgs::Log>()},
    {"rosgraph_msgs/TopicStatistics", new EchoSubscriber<rosgraph_msgs::TopicStatistics>()},
    {"sensor_msgs/BatteryState", new EchoSubscriber<sensor_msgs::BatteryState>()},
    {"sensor_msgs/CameraInfo", new EchoSubscriber<sensor_msgs::CameraInfo>()},
    {"sensor_msgs/ChannelFloat32", new EchoSubscriber<sensor_msgs::ChannelFloat32>()},
    {"sensor_msgs/CompressedImage", new EchoSubscriber<sensor_msgs::CompressedImage>()},
    {"sensor_msgs/FluidPressure", new EchoSubscriber<sensor_msgs::FluidPressure>()},
    {"sensor_msgs/Illuminance", new EchoSubscriber<sensor_msgs::Illuminance>()},
    {"sensor_msgs/Image", new EchoSubscriber<sensor_msgs::Image>()},
    {"sensor_msgs/Imu", new EchoSubscriber<sensor_msgs::Imu>()},
    {"sensor_msgs/JointState", new EchoSubscriber<sensor_msgs::JointState>()},
    {"sensor_msgs/Joy", new EchoSubscriber<sensor_msgs::Joy>()},
    {"sensor_msgs/JoyFeedback", new EchoSubscriber<sensor_msgs::JoyFeedback>()},
    {"sensor_msgs/JoyFeedbackArray", new EchoSubscriber<sensor_msgs::JoyFeedbackArray>()},
    {"sensor_msgs/LaserEcho", new EchoSubscriber<sensor_msgs::LaserEcho>()},
    {"sensor_msgs/LaserScan", new EchoSubscriber<sensor_msgs::LaserScan>()},
    {"sensor_msgs/MagneticField", new EchoSubscriber<sensor_msgs::MagneticField>()},
    {"sensor_msgs/MultiDOFJointState", new EchoSubscriber<sensor_msgs::MultiDOFJointState>()},
    {"sensor_msgs/MultiEchoLaserScan", new EchoSubscriber<sensor_msgs::MultiEchoLaserScan>()},
    {"sensor_msgs/NavSatFix", new EchoSubscriber<sensor_msgs::NavSatFix>()},
    {"sensor_msgs/NavSatStatus", new EchoSubscriber<sensor_msgs::NavSatStatus>()},
    {"sensor_msgs/PointCloud", new EchoSubscriber<sensor_msgs::PointCloud>()},
    {"sensor_msgs/PointCloud2", new EchoSubscriber<sensor_msgs::PointCloud2>()},
    {"sensor_msgs/PointField", new EchoSubscriber<sensor_msgs::PointField>()},
    {"sensor_msgs/Range", new EchoSubscriber<sensor_msgs::Range>()},
    {"sensor_msgs/RegionOfInterest", new EchoSubscriber<sensor_msgs::RegionOfInterest>()},
    {"sensor_msgs/RelativeHumidity", new EchoSubscriber<sensor_msgs::RelativeHumidity>()},
    {"sensor_msgs/Temperature", new EchoSubscriber<sensor_msgs::Temperature>()},
    {"sensor_msgs/TimeReference", new EchoSubscriber<sensor_msgs::TimeReference>()},
    {"shape_msgs/Mesh", new EchoSubscriber<shape_msgs::Mesh>()},
    {"shape_msgs/MeshTriangle", new EchoSubscriber<shape_msgs::MeshTriangle>()},
    {"shape_msgs/Plane", new EchoSubscriber<shape_msgs::Plane>()},
    {"shape_msgs/SolidPrimitive", new EchoSubscriber<shape_msgs::SolidPrimitive>()},
    {"smach_msgs/SmachContainerInitialStatusCmd", new EchoSubscriber<smach_msgs::SmachContainerInitialStatusCmd>()},
    {"smach_msgs/SmachContainerStatus", new EchoSubscriber<smach_msgs::SmachContainerStatus>()},
    {"smach_msgs/SmachContainerStructure", new EchoSubscriber<smach_msgs::SmachContainerStructure>()},
    {"std_msgs/Bool", new EchoSubscriber<std_msgs::Bool>()},
    {"std_msgs/Byte", new EchoSubscriber<std_msgs::Byte>()},
    {"std_msgs/ByteMultiArray", new EchoSubscriber<std_msgs::ByteMultiArray>()},
    {"std_msgs/Char", new EchoSubscriber<std_msgs::Char>()},
    {"std_msgs/ColorRGBA", new EchoSubscriber<std_msgs::ColorRGBA>()},
    {"std_msgs/Duration", new EchoSubscriber<std_msgs::Duration>()},
    {"std_msgs/Empty", new EchoSubscriber<std_msgs::Empty>()},
    {"std_msgs/Float32", new EchoSubscriber<std_msgs::Float32>()},
    {"std_msgs/Float32MultiArray", new EchoSubscriber<std_msgs::Float32MultiArray>()},
    {"std_msgs/Float64", new EchoSubscriber<std_msgs::Float64>()},
    {"std_msgs/Float64MultiArray", new EchoSubscriber<std_msgs::Float64MultiArray>()},
    {"std_msgs/Header", new EchoSubscriber<std_msgs::Header>()},
    {"std_msgs/Int16", new EchoSubscriber<std_msgs::Int16>()},
    {"std_msgs/Int16MultiArray", new EchoSubscriber<std_msgs::Int16MultiArray>()},
    {"std_msgs/Int32", new EchoSubscriber<std_msgs::Int32>()},
    {"std_msgs/Int32MultiArray", new EchoSubscriber<std_msgs::Int32MultiArray>()},
    {"std_msgs/Int64", new EchoSubscriber<std_msgs::Int64>()},
    {"std_msgs/Int64MultiArray", new EchoSubscriber<std_msgs::Int64MultiArray>()},
    {"std_msgs/Int8", new EchoSubscriber<std_msgs::Int8>()},
    {"std_msgs/Int8MultiArray", new EchoSubscriber<std_msgs::Int8MultiArray>()},
    {"std_msgs/MultiArrayDimension", new EchoSubscriber<std_msgs::MultiArrayDimension>()},
    {"std_msgs/MultiArrayLayout", new EchoSubscriber<std_msgs::MultiArrayLayout>()},
    {"std_msgs/String", new EchoSubscriber<std_msgs::String>()},
    {"std_msgs/Time", new EchoSubscriber<std_msgs::Time>()},
    {"std_msgs/UInt16", new EchoSubscriber<std_msgs::UInt16>()},
    {"std_msgs/UInt16MultiArray", new EchoSubscriber<std_msgs::UInt16MultiArray>()},
    {"std_msgs/UInt32", new EchoSubscriber<std_msgs::UInt32>()},
    {"std_msgs/UInt32MultiArray", new EchoSubscriber<std_msgs::UInt32MultiArray>()},
    {"std_msgs/UInt64", new EchoSubscriber<std_msgs::UInt64>()},
    {"std_msgs/UInt64MultiArray", new EchoSubscriber<std_msgs::UInt64MultiArray>()},
    {"std_msgs/UInt8", new EchoSubscriber<std_msgs::UInt8>()},
    {"std_msgs/UInt8MultiArray", new EchoSubscriber<std_msgs::UInt8MultiArray>()},
    {"stereo_msgs/DisparityImage", new EchoSubscriber<stereo_msgs::DisparityImage>()},
    {"tf/tfMessage", new EchoSubscriber<tf::tfMessage>()},
    {"tf2_msgs/LookupTransformAction", new EchoSubscriber<tf2_msgs::LookupTransformAction>()},
    {"tf2_msgs/LookupTransformActionFeedback", new EchoSubscriber<tf2_msgs::LookupTransformActionFeedback>()},
    {"tf2_msgs/LookupTransformActionGoal", new EchoSubscriber<tf2_msgs::LookupTransformActionGoal>()},
    {"tf2_msgs/LookupTransformActionResult", new EchoSubscriber<tf2_msgs::LookupTransformActionResult>()},
    {"tf2_msgs/LookupTransformFeedback", new EchoSubscriber<tf2_msgs::LookupTransformFeedback>()},
    {"tf2_msgs/LookupTransformGoal", new EchoSubscriber<tf2_msgs::LookupTransformGoal>()},
    {"tf2_msgs/LookupTransformResult", new EchoSubscriber<tf2_msgs::LookupTransformResult>()},
    {"tf2_msgs/TF2Error", new EchoSubscriber<tf2_msgs::TF2Error>()},
    {"tf2_msgs/TFMessage", new EchoSubscriber<tf2_msgs::TFMessage>()},
    {"tinyros_hello/TinyrosHello", new EchoSubscriber<tinyros_hello::TinyrosHello>()},
    {"tinyros_msgs/Log", new EchoSubscriber<tinyros_msgs::Log>()},
    {"tinyros_msgs/TopicInfo", new EchoSubscriber<tinyros_msgs::TopicInfo>()},
    {"trajectory_msgs/JointTrajectory", new EchoSubscriber<trajectory_msgs::JointTrajectory>()},
    {"trajectory_msgs/JointTrajectoryPoint", new EchoSubscriber<trajectory_msgs::JointTrajectoryPoint>()},
    {"trajectory_msgs/MultiDOFJointTrajectory", new EchoSubscriber<trajectory_msgs::MultiDOFJointTrajectory>()},
    {"trajectory_msgs/MultiDOFJointTrajectoryPoint", new EchoSubscriber<trajectory_msgs::MultiDOFJointTrajectoryPoint>()},
};

}
#endif

