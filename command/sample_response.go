package command

// The sample response for tests.

var srDescribeDBInstanceResponse = `
<DescribeDBInstancesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBInstancesResult>
    <DBInstances>
      <DBInstance>
        <BackupRetentionPeriod>0</BackupRetentionPeriod>
        <DBInstanceStatus>available</DBInstanceStatus>
        <MultiAZ>false</MultiAZ>
        <VpcSecurityGroups>
          <VpcSecurityGroupMembership>
            <Status>active</Status>
            <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
          </VpcSecurityGroupMembership>
        </VpcSecurityGroups>
        <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
        <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
        <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
        <AvailabilityZone>us-west-2c</AvailabilityZone>
        <ReadReplicaDBInstanceIdentifiers/>
        <Engine>mysql</Engine>
        <PendingModifiedValues/>
        <LicenseModel>general-public-license</LicenseModel>
        <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
        <DBSubnetGroup>
          <VpcId>vpc-1a12bc34</VpcId>
          <SubnetGroupStatus>Complete</SubnetGroupStatus>
          <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
          <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
          <Subnets>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2c</Name>
              </SubnetAvailabilityZone>
            </Subnet>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2a</Name>
              </SubnetAvailabilityZone>
            </Subnet>
          </Subnets>
        </DBSubnetGroup>
        <EngineVersion>5.6.13</EngineVersion>
        <DBParameterGroups>
          <DBParameterGroup>
            <ParameterApplyStatus>in-sync</ParameterApplyStatus>
            <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
          </DBParameterGroup>
        </DBParameterGroups>
        <Endpoint>
          <Port>3306</Port>
          <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
        </Endpoint>
        <OptionGroupMemberships>
          <OptionGroupMembership>
            <OptionGroupName>default:mysql-5-6</OptionGroupName>
            <Status>in-sync</Status>
          </OptionGroupMembership>
        </OptionGroupMemberships>
        <DBSecurityGroups/>
        <PubliclyAccessible>false</PubliclyAccessible>
        <DBName>rdstrytestdb</DBName>
        <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
        <StorageType>gp2</StorageType>
        <InstanceCreateTime>2015-02-18T02:03:37.927Z</InstanceCreateTime>
        <AllocatedStorage>5</AllocatedStorage>
        <StorageEncrypted>false</StorageEncrypted>
        <MasterUsername>testroot</MasterUsername>
        <DBInstanceClass>db.t1.micro</DBInstanceClass>
      </DBInstance>
    </DBInstances>
  </DescribeDBInstancesResult>
  <ResponseMetadata>
    <RequestId>9cbcf47c-b716-11e4-abc4-cd23b33bbeed</RequestId>
  </ResponseMetadata>
</DescribeDBInstancesResponse>
`
var srListTagsForResourceResponse = `
<ListTagsForResourceResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ListTagsForResourceResult>
    <TagList>
      <Tag>
        <Value>2015-02-18-10-59-05</Value>
        <Key>rt_time</Key>
      </Tag>
      <Tag>
        <Value>rds-try-v0-0-1</Value>
        <Key>rt_name</Key>
      </Tag>
    </TagList>
  </ListTagsForResourceResult>
  <ResponseMetadata>
    <RequestId>f6d500fb-b718-11e4-a490-772d0a37354b</RequestId>
  </ResponseMetadata>
</ListTagsForResourceResponse>
`
var srDescribeDBInstancesResponse = `
<DescribeDBInstancesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBInstancesResult>
    <DBInstances>
      <DBInstance>
        <BackupRetentionPeriod>0</BackupRetentionPeriod>
        <DBInstanceStatus>available</DBInstanceStatus>
        <MultiAZ>false</MultiAZ>
        <VpcSecurityGroups>
          <VpcSecurityGroupMembership>
            <Status>active</Status>
            <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
          </VpcSecurityGroupMembership>
        </VpcSecurityGroups>
        <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
        <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
        <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
        <AvailabilityZone>us-west-2c</AvailabilityZone>
        <ReadReplicaDBInstanceIdentifiers/>
        <Engine>mysql</Engine>
        <PendingModifiedValues/>
        <LicenseModel>general-public-license</LicenseModel>
        <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
        <DBSubnetGroup>
          <VpcId>vpc-1a12bc34</VpcId>
          <SubnetGroupStatus>Complete</SubnetGroupStatus>
          <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
          <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
          <Subnets>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2c</Name>
              </SubnetAvailabilityZone>
            </Subnet>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2a</Name>
              </SubnetAvailabilityZone>
            </Subnet>
          </Subnets>
        </DBSubnetGroup>
        <EngineVersion>5.6.13</EngineVersion>
        <DBParameterGroups>
          <DBParameterGroup>
            <ParameterApplyStatus>in-sync</ParameterApplyStatus>
            <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
          </DBParameterGroup>
        </DBParameterGroups>
        <Endpoint>
          <Port>3306</Port>
          <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
        </Endpoint>
        <OptionGroupMemberships>
          <OptionGroupMembership>
            <OptionGroupName>default:mysql-5-6</OptionGroupName>
            <Status>in-sync</Status>
          </OptionGroupMembership>
        </OptionGroupMemberships>
        <DBSecurityGroups/>
        <PubliclyAccessible>false</PubliclyAccessible>
        <DBName>rdstrytestdb</DBName>
        <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
        <StorageType>gp2</StorageType>
        <InstanceCreateTime>2015-02-18T02:03:37.927Z</InstanceCreateTime>
        <AllocatedStorage>5</AllocatedStorage>
        <StorageEncrypted>false</StorageEncrypted>
        <MasterUsername>testroot</MasterUsername>
        <DBInstanceClass>db.t1.micro</DBInstanceClass>
      </DBInstance>
      <DBInstance>
        <BackupRetentionPeriod>0</BackupRetentionPeriod>
        <DBInstanceStatus>available</DBInstanceStatus>
        <MultiAZ>false</MultiAZ>
        <VpcSecurityGroups>
          <VpcSecurityGroupMembership>
            <Status>active</Status>
            <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
          </VpcSecurityGroupMembership>
        </VpcSecurityGroups>
        <DBInstanceIdentifier>rds-try-test-db-2</DBInstanceIdentifier>
        <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
        <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
        <AvailabilityZone>us-west-2c</AvailabilityZone>
        <ReadReplicaDBInstanceIdentifiers/>
        <Engine>mysql</Engine>
        <PendingModifiedValues/>
        <LicenseModel>general-public-license</LicenseModel>
        <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
        <DBSubnetGroup>
          <VpcId>vpc-1a12bc34</VpcId>
          <SubnetGroupStatus>Complete</SubnetGroupStatus>
          <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
          <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
          <Subnets>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2c</Name>
              </SubnetAvailabilityZone>
            </Subnet>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2a</Name>
              </SubnetAvailabilityZone>
            </Subnet>
          </Subnets>
        </DBSubnetGroup>
        <EngineVersion>5.6.13</EngineVersion>
        <DBParameterGroups>
          <DBParameterGroup>
            <ParameterApplyStatus>in-sync</ParameterApplyStatus>
            <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
          </DBParameterGroup>
        </DBParameterGroups>
        <Endpoint>
          <Port>3306</Port>
          <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
        </Endpoint>
        <OptionGroupMemberships>
          <OptionGroupMembership>
            <OptionGroupName>default:mysql-5-6</OptionGroupName>
            <Status>in-sync</Status>
          </OptionGroupMembership>
        </OptionGroupMemberships>
        <DBSecurityGroups/>
        <PubliclyAccessible>false</PubliclyAccessible>
        <DBName>rdstrytestdb</DBName>
        <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
        <StorageType>gp2</StorageType>
        <InstanceCreateTime>2015-02-18T02:27:41.554Z</InstanceCreateTime>
        <AllocatedStorage>5</AllocatedStorage>
        <StorageEncrypted>false</StorageEncrypted>
        <MasterUsername>testroot</MasterUsername>
        <DBInstanceClass>db.t1.micro</DBInstanceClass>
      </DBInstance>
      <DBInstance>
        <BackupRetentionPeriod>0</BackupRetentionPeriod>
        <DBInstanceStatus>available</DBInstanceStatus>
        <MultiAZ>false</MultiAZ>
        <VpcSecurityGroups>
          <VpcSecurityGroupMembership>
            <Status>active</Status>
            <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
          </VpcSecurityGroupMembership>
        </VpcSecurityGroups>
        <DBInstanceIdentifier>rds-try-test-db-3</DBInstanceIdentifier>
        <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
        <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
        <AvailabilityZone>us-west-2a</AvailabilityZone>
        <ReadReplicaDBInstanceIdentifiers/>
        <Engine>mysql</Engine>
        <PendingModifiedValues/>
        <LicenseModel>general-public-license</LicenseModel>
        <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
        <DBSubnetGroup>
          <VpcId>vpc-1a12bc34</VpcId>
          <SubnetGroupStatus>Complete</SubnetGroupStatus>
          <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
          <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
          <Subnets>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2c</Name>
              </SubnetAvailabilityZone>
            </Subnet>
            <Subnet>
              <SubnetStatus>Active</SubnetStatus>
              <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
              <SubnetAvailabilityZone>
                <Name>us-west-2a</Name>
              </SubnetAvailabilityZone>
            </Subnet>
          </Subnets>
        </DBSubnetGroup>
        <EngineVersion>5.6.13</EngineVersion>
        <DBParameterGroups>
          <DBParameterGroup>
            <ParameterApplyStatus>in-sync</ParameterApplyStatus>
            <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
          </DBParameterGroup>
        </DBParameterGroups>
        <Endpoint>
          <Port>3306</Port>
          <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
        </Endpoint>
        <OptionGroupMemberships>
          <OptionGroupMembership>
            <OptionGroupName>default:mysql-5-6</OptionGroupName>
            <Status>in-sync</Status>
          </OptionGroupMembership>
        </OptionGroupMemberships>
        <DBSecurityGroups/>
        <PubliclyAccessible>false</PubliclyAccessible>
        <DBName>rdstrytestdb</DBName>
        <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
        <StorageType>gp2</StorageType>
        <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
        <AllocatedStorage>5</AllocatedStorage>
        <StorageEncrypted>false</StorageEncrypted>
        <MasterUsername>testroot</MasterUsername>
        <DBInstanceClass>db.t1.micro</DBInstanceClass>
      </DBInstance>
    </DBInstances>
  </DescribeDBInstancesResult>
  <ResponseMetadata>
    <RequestId>9cbcf47c-b716-11e4-abc4-cd23b33bbeed</RequestId>
  </ResponseMetadata>
</DescribeDBInstancesResponse>
`
var srModifyDBInstanceResponse = `
<ModifyDBInstanceResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ModifyDBInstanceResult>
    <DBInstance>
      <BackupRetentionPeriod>0</BackupRetentionPeriod>
      <DBInstanceStatus>available</DBInstanceStatus>
      <MultiAZ>false</MultiAZ>
      <VpcSecurityGroups>
        <VpcSecurityGroupMembership>
          <Status>adding</Status>
          <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
        </VpcSecurityGroupMembership>
        <VpcSecurityGroupMembership>
          <Status>removing</Status>
          <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
        </VpcSecurityGroupMembership>
      </VpcSecurityGroups>
      <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
      <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
      <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
      <AvailabilityZone>us-west-2c</AvailabilityZone>
      <ReadReplicaDBInstanceIdentifiers/>
      <Engine>mysql</Engine>
      <PendingModifiedValues/>
      <LicenseModel>general-public-license</LicenseModel>
      <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
      <DBSubnetGroup>
        <VpcId>vpc-1a12bc34</VpcId>
        <SubnetGroupStatus>Complete</SubnetGroupStatus>
        <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
        <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
        <Subnets>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2c</Name>
            </SubnetAvailabilityZone>
          </Subnet>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2a</Name>
            </SubnetAvailabilityZone>
          </Subnet>
        </Subnets>
      </DBSubnetGroup>
      <EngineVersion>5.6.13</EngineVersion>
      <DBParameterGroups>
        <DBParameterGroup>
          <ParameterApplyStatus>applying</ParameterApplyStatus>
          <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
        </DBParameterGroup>
      </DBParameterGroups>
      <Endpoint>
        <Port>3306</Port>
        <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
      </Endpoint>
      <OptionGroupMemberships>
        <OptionGroupMembership>
          <OptionGroupName>default:mysql-5-6</OptionGroupName>
          <Status>in-sync</Status>
        </OptionGroupMembership>
      </OptionGroupMemberships>
      <DBSecurityGroups/>
      <PubliclyAccessible>false</PubliclyAccessible>
      <DBName>rdstrytestdb</DBName>
      <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
      <StorageType>gp2</StorageType>
      <InstanceCreateTime>2015-02-18T02:41:49.166Z</InstanceCreateTime>
      <AllocatedStorage>5</AllocatedStorage>
      <StorageEncrypted>false</StorageEncrypted>
      <MasterUsername>testroot</MasterUsername>
      <DBInstanceClass>db.t1.micro</DBInstanceClass>
    </DBInstance>
  </ModifyDBInstanceResult>
  <ResponseMetadata>
    <RequestId>e590570f-b717-11e4-abc4-cd23b33bbeed</RequestId>
  </ResponseMetadata>
</ModifyDBInstanceResponse>
`
var srRebootDBInstanceResponse = `
<RebootDBInstanceResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <RebootDBInstanceResult>
    <DBInstance>
      <BackupRetentionPeriod>0</BackupRetentionPeriod>
      <MultiAZ>false</MultiAZ>
      <DBInstanceStatus>rebooting</DBInstanceStatus>
      <VpcSecurityGroups>
        <VpcSecurityGroupMembership>
          <Status>active</Status>
          <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
        </VpcSecurityGroupMembership>
      </VpcSecurityGroups>
      <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
      <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
      <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
      <AvailabilityZone>us-west-2c</AvailabilityZone>
      <ReadReplicaDBInstanceIdentifiers/>
      <Engine>mysql</Engine>
      <PendingModifiedValues/>
      <LicenseModel>general-public-license</LicenseModel>
      <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
      <DBSubnetGroup>
        <VpcId>vpc-1a12bc34</VpcId>
        <SubnetGroupStatus>Complete</SubnetGroupStatus>
        <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
        <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
        <Subnets>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2c</Name>
            </SubnetAvailabilityZone>
          </Subnet>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2a</Name>
            </SubnetAvailabilityZone>
          </Subnet>
        </Subnets>
      </DBSubnetGroup>
      <Endpoint>
        <Port>3306</Port>
        <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
      </Endpoint>
      <EngineVersion>5.6.13</EngineVersion>
      <DBParameterGroups>
        <DBParameterGroup>
          <ParameterApplyStatus>applying</ParameterApplyStatus>
          <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
        </DBParameterGroup>
      </DBParameterGroups>
      <OptionGroupMemberships>
        <OptionGroupMembership>
          <OptionGroupName>default:mysql-5-6</OptionGroupName>
          <Status>in-sync</Status>
        </OptionGroupMembership>
      </OptionGroupMemberships>
      <DBSecurityGroups/>
      <PubliclyAccessible>false</PubliclyAccessible>
      <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
      <DBName>rdstrytestdb</DBName>
      <StorageType>gp2</StorageType>
      <InstanceCreateTime>2015-02-18T02:41:49.166Z</InstanceCreateTime>
      <AllocatedStorage>5</AllocatedStorage>
      <StorageEncrypted>false</StorageEncrypted>
      <DBInstanceClass>db.t1.micro</DBInstanceClass>
      <MasterUsername>testroot</MasterUsername>
    </DBInstance>
  </RebootDBInstanceResult>
  <ResponseMetadata>
    <RequestId>f7f18a44-b717-11e4-873c-218142f1d71d</RequestId>
  </ResponseMetadata>
</RebootDBInstanceResponse>
`
var srDescribeDBSnapshotsResponse = `
<DescribeDBSnapshotsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBSnapshotsResult>
    <DBSnapshots>
      <DBSnapshot>
        <Port>3306</Port>
        <OptionGroupName>default:mysql-5-6</OptionGroupName>
        <Engine>mysql</Engine>
        <Status>available</Status>
        <SnapshotType>manual</SnapshotType>
        <LicenseModel>general-public-license</LicenseModel>
        <EngineVersion>5.6.13</EngineVersion>
        <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
        <DBSnapshotIdentifier>before-test-1</DBSnapshotIdentifier>
        <SnapshotCreateTime>2014-08-25T10:11:10.622Z</SnapshotCreateTime>
        <Encrypted>false</Encrypted>
        <VpcId>vpc-1a12bc34</VpcId>
        <StorageType>standard</StorageType>
        <AvailabilityZone>us-west-2a</AvailabilityZone>
        <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
        <PercentProgress>100</PercentProgress>
        <AllocatedStorage>5</AllocatedStorage>
        <MasterUsername>testroot</MasterUsername>
      </DBSnapshot>
      <DBSnapshot>
        <Port>3306</Port>
        <OptionGroupName>default:mysql-5-6</OptionGroupName>
        <Engine>mysql</Engine>
        <Status>available</Status>
        <SnapshotType>manual</SnapshotType>
        <LicenseModel>general-public-license</LicenseModel>
        <EngineVersion>5.6.13</EngineVersion>
        <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
        <DBSnapshotIdentifier>before-test-2</DBSnapshotIdentifier>
        <SnapshotCreateTime>2014-09-05T05:36:00.675Z</SnapshotCreateTime>
        <Encrypted>false</Encrypted>
        <VpcId>vpc-1a12bc34</VpcId>
        <StorageType>standard</StorageType>
        <AvailabilityZone>us-west-2a</AvailabilityZone>
        <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
        <PercentProgress>100</PercentProgress>
        <AllocatedStorage>5</AllocatedStorage>
        <MasterUsername>testroot</MasterUsername>
      </DBSnapshot>
      <DBSnapshot>
        <Port>3306</Port>
        <OptionGroupName>default:mysql-5-6</OptionGroupName>
        <Engine>mysql</Engine>
        <Status>available</Status>
        <SnapshotType>manual</SnapshotType>
        <LicenseModel>general-public-license</LicenseModel>
        <EngineVersion>5.6.13</EngineVersion>
        <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
        <DBSnapshotIdentifier>before-test-3</DBSnapshotIdentifier>
        <SnapshotCreateTime>2014-03-12T02:37:30.169Z</SnapshotCreateTime>
        <Encrypted>false</Encrypted>
        <VpcId>vpc-1a12bc34</VpcId>
        <StorageType>standard</StorageType>
        <AvailabilityZone>us-west-2a</AvailabilityZone>
        <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
        <PercentProgress>100</PercentProgress>
        <AllocatedStorage>5</AllocatedStorage>
        <MasterUsername>testroot</MasterUsername>
      </DBSnapshot>
    </DBSnapshots>
  </DescribeDBSnapshotsResult>
  <ResponseMetadata>
    <RequestId>a118adf4-b716-11e4-873c-218142f1d71d</RequestId>
  </ResponseMetadata>
</DescribeDBSnapshotsResponse>
`
var srDescribeDBSnapshotResponse = `
<DescribeDBSnapshotsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBSnapshotsResult>
    <DBSnapshots>
      <DBSnapshot>
        <Port>3306</Port>
        <OptionGroupName>default:mysql-5-6</OptionGroupName>
        <Engine>mysql</Engine>
        <Status>available</Status>
        <SnapshotType>manual</SnapshotType>
        <LicenseModel>general-public-license</LicenseModel>
        <EngineVersion>5.6.13</EngineVersion>
        <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
        <DBSnapshotIdentifier>before-test-1</DBSnapshotIdentifier>
        <SnapshotCreateTime>2014-08-25T10:11:10.622Z</SnapshotCreateTime>
        <Encrypted>false</Encrypted>
        <VpcId>vpc-1a12bc34</VpcId>
        <StorageType>standard</StorageType>
        <AvailabilityZone>us-west-2a</AvailabilityZone>
        <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
        <PercentProgress>100</PercentProgress>
        <AllocatedStorage>5</AllocatedStorage>
        <MasterUsername>testroot</MasterUsername>
      </DBSnapshot>
    </DBSnapshots>
  </DescribeDBSnapshotsResult>
  <ResponseMetadata>
    <RequestId>a118adf4-b716-11e4-873c-218142f1d71d</RequestId>
  </ResponseMetadata>
</DescribeDBSnapshotsResponse>
`
var srDeleteDBInstanceResponse = `
<DeleteDBInstanceResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DeleteDBInstanceResult>
    <DBInstance>
      <BackupRetentionPeriod>0</BackupRetentionPeriod>
      <DBInstanceStatus>deleting</DBInstanceStatus>
      <MultiAZ>false</MultiAZ>
      <VpcSecurityGroups>
        <VpcSecurityGroupMembership>
          <Status>active</Status>
          <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
        </VpcSecurityGroupMembership>
      </VpcSecurityGroups>
      <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
      <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
      <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
      <AvailabilityZone>us-west-2c</AvailabilityZone>
      <ReadReplicaDBInstanceIdentifiers/>
      <Engine>mysql</Engine>
      <PendingModifiedValues/>
      <LicenseModel>general-public-license</LicenseModel>
      <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
      <DBSubnetGroup>
        <VpcId>vpc-1a12bc34</VpcId>
        <SubnetGroupStatus>Complete</SubnetGroupStatus>
        <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
        <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
        <Subnets>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2c</Name>
            </SubnetAvailabilityZone>
          </Subnet>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2a</Name>
            </SubnetAvailabilityZone>
          </Subnet>
        </Subnets>
      </DBSubnetGroup>
      <EngineVersion>5.6.13</EngineVersion>
      <DBParameterGroups>
        <DBParameterGroup>
          <ParameterApplyStatus>in-sync</ParameterApplyStatus>
          <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
        </DBParameterGroup>
      </DBParameterGroups>
      <Endpoint>
        <Port>3306</Port>
        <Address>rds-try-test-db.c6c2mntzugv0.us-west-2.rds.amazonaws.com</Address>
      </Endpoint>
      <OptionGroupMemberships>
        <OptionGroupMembership>
          <OptionGroupName>default:mysql-5-6</OptionGroupName>
          <Status>in-sync</Status>
        </OptionGroupMembership>
      </OptionGroupMemberships>
      <DBSecurityGroups/>
      <PubliclyAccessible>false</PubliclyAccessible>
      <DBName>rdstrytestdb</DBName>
      <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
      <StorageType>gp2</StorageType>
      <InstanceCreateTime>2015-02-18T02:03:37.927Z</InstanceCreateTime>
      <AllocatedStorage>5</AllocatedStorage>
      <StorageEncrypted>false</StorageEncrypted>
      <MasterUsername>testroot</MasterUsername>
      <DBInstanceClass>db.t1.micro</DBInstanceClass>
    </DBInstance>
  </DeleteDBInstanceResult>
  <ResponseMetadata>
    <RequestId>f81dadbd-b718-11e4-8115-39659aea07b9</RequestId>
  </ResponseMetadata>
</DeleteDBInstanceResponse>
`
var srDeleteDBSnapshotResponse = `
<DeleteDBSnapshotResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DeleteDBSnapshotResult>
    <DBSnapshot>
      <Port>3306</Port>
      <OptionGroupName>default:mysql-5-6</OptionGroupName>
      <Engine>mysql</Engine>
      <Status>deleted</Status>
      <SnapshotType>manual</SnapshotType>
      <LicenseModel>general-public-license</LicenseModel>
      <EngineVersion>5.6.13</EngineVersion>
      <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
      <DBSnapshotIdentifier>before-test-1</DBSnapshotIdentifier>
      <SnapshotCreateTime>2015-02-18T01:57:25.758Z</SnapshotCreateTime>
      <Encrypted>false</Encrypted>
      <VpcId>vpc-1a12bc34</VpcId>
      <StorageType>gp2</StorageType>
      <AvailabilityZone>us-west-2a</AvailabilityZone>
      <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
      <PercentProgress>100</PercentProgress>
      <AllocatedStorage>5</AllocatedStorage>
      <MasterUsername>root</MasterUsername>
    </DBSnapshot>
  </DeleteDBSnapshotResult>
  <ResponseMetadata>
    <RequestId>f8b81909-b718-11e4-8115-39659aea07b9</RequestId>
  </ResponseMetadata>
</DeleteDBSnapshotResponse>
`
var srCreateDBSnapshotResponse = `
<CreateDBSnapshotResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateDBSnapshotResult>
    <DBSnapshot>
      <Port>3306</Port>
      <OptionGroupName>default:mysql-5-6</OptionGroupName>
      <Status>creating</Status>
      <Engine>mysql</Engine>
      <SnapshotType>manual</SnapshotType>
      <LicenseModel>general-public-license</LicenseModel>
      <EngineVersion>5.6.13</EngineVersion>
      <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
      <DBSnapshotIdentifier>before-test-1</DBSnapshotIdentifier>
      <Encrypted>false</Encrypted>
      <VpcId>vpc-1a12bc34</VpcId>
      <StorageType>gp2</StorageType>
      <AvailabilityZone>us-west-2a</AvailabilityZone>
      <InstanceCreateTime>2014-05-16T11:30:16.183Z</InstanceCreateTime>
      <PercentProgress>0</PercentProgress>
      <AllocatedStorage>5</AllocatedStorage>
      <MasterUsername>testroot</MasterUsername>
    </DBSnapshot>
  </CreateDBSnapshotResult>
  <ResponseMetadata>
    <RequestId>a136480a-b716-11e4-abc4-cd23b33bbeed</RequestId>
  </ResponseMetadata>
</CreateDBSnapshotResponse>
`
var srRestoreDBInstanceFromDBSnapshotResponse = `
<RestoreDBInstanceFromDBSnapshotResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <RestoreDBInstanceFromDBSnapshotResult>
    <DBInstance>
      <BackupRetentionPeriod>0</BackupRetentionPeriod>
      <MultiAZ>false</MultiAZ>
      <DBInstanceStatus>creating</DBInstanceStatus>
      <VpcSecurityGroups>
        <VpcSecurityGroupMembership>
          <Status>active</Status>
          <VpcSecurityGroupId>sg-123a456b</VpcSecurityGroupId>
        </VpcSecurityGroupMembership>
      </VpcSecurityGroups>
      <DBInstanceIdentifier>rds-try-test-db-1</DBInstanceIdentifier>
      <PreferredBackupWindow>18:00-18:30</PreferredBackupWindow>
      <PreferredMaintenanceWindow>wed:19:00-wed:19:30</PreferredMaintenanceWindow>
      <ReadReplicaDBInstanceIdentifiers/>
      <Engine>mysql</Engine>
      <PendingModifiedValues/>
      <LicenseModel>general-public-license</LicenseModel>
      <DbiResourceId>db-ABCDEFGHIJKLNMOPQRSTUVWXYZ</DbiResourceId>
      <DBSubnetGroup>
        <VpcId>vpc-1a12bc34</VpcId>
        <SubnetGroupStatus>Complete</SubnetGroupStatus>
        <DBSubnetGroupDescription>mysql-subnet-group</DBSubnetGroupDescription>
        <DBSubnetGroupName>mysql-subnet-group</DBSubnetGroupName>
        <Subnets>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a56b</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2c</Name>
            </SubnetAvailabilityZone>
          </Subnet>
          <Subnet>
            <SubnetStatus>Active</SubnetStatus>
            <SubnetIdentifier>subnet-1234a5bc</SubnetIdentifier>
            <SubnetAvailabilityZone>
              <Name>us-west-2a</Name>
            </SubnetAvailabilityZone>
          </Subnet>
        </Subnets>
      </DBSubnetGroup>
      <EngineVersion>5.6.13</EngineVersion>
      <DBParameterGroups>
        <DBParameterGroup>
          <ParameterApplyStatus>in-sync</ParameterApplyStatus>
          <DBParameterGroupName>default:mysql-5-6</DBParameterGroupName>
        </DBParameterGroup>
      </DBParameterGroups>
      <OptionGroupMemberships>
        <OptionGroupMembership>
          <OptionGroupName>default:mysql-5-6</OptionGroupName>
          <Status>pending-apply</Status>
        </OptionGroupMembership>
      </OptionGroupMemberships>
      <DBSecurityGroups/>
      <PubliclyAccessible>false</PubliclyAccessible>
      <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
      <DBName>rdstrytestdb</DBName>
      <StorageType>gp2</StorageType>
      <AllocatedStorage>5</AllocatedStorage>
      <StorageEncrypted>false</StorageEncrypted>
      <DBInstanceClass>db.t1.micro</DBInstanceClass>
      <MasterUsername>testroot</MasterUsername>
    </DBInstance>
  </RestoreDBInstanceFromDBSnapshotResult>
  <ResponseMetadata>
    <RequestId>e96f3397-b716-11e4-873c-218142f1d71d</RequestId>
  </ResponseMetadata>
</RestoreDBInstanceFromDBSnapshotResponse>
`
