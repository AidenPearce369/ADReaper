# Proof Of Concept

Working PoC of ```ADReaper```

To get the list of all supported commands by ```ADReaper```,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe

      -command string

            Command to run
                  dc              - to list domain controllers
                  domain-trust    - to list domain trust
                  users           - to list all users
                  computers       - to list all computers
                  groups          - to list all groups with members
                  spn             - to list service principal objects
                  never-loggedon  - to list users never logged on
                  gpo             - to list group policy objects
                  ou              - to list organizational units
                  ms-sql          - to list MS-SQL servers
                  asreproast      - to list AS-REP roastable accounts
                  unconstrained   - to list Unconstrained Delegated accounts
                  admin-priv      - to list AD objects with admin privilege

      -dc string

            Enter the DC

      -filter string

            Filters to use for users/groups/computers

            list - lists all objects only
            fulldata - list all objects with properties
            membership - lists all members from an object

            (default "list")
      -name string

            Pass object name of user/group/computer

      -password string

            Enter the Password

      -user string

            Enter the Username
```

To query the attributes of ```Domain Controller```,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command dc
+------------------------------------------------------------------------------------------------+
| DOMAIN CONTROLLERS                                                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN=UFC-DC1,OU=Domain Controllers,DC=us,DC=funcorp,DC=local                                |
| SAM Account Name : UFC-DC1$                                                                    |
| SAM Account Type : 805306369                                                                   |
| CN : UFC-DC1                                                                                   |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user, computer                               |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local                   |
| DNS Host Name : rt-dc.rt.securecompany.com                                                       |
| Service Principal Name (SPN) :                                                                 |
|     ldap/rt-dc.rt.securecompany.com/DomainDnsZones.rt.securecompany.com                              |
|     ldap/rt-dc.rt.securecompany.com/ForestDnsZones.funcorp.local                                 |
|     Dfsr-12F9A27C-BF97-4787-9364-D31B6C55EB04/rt-dc.rt.securecompany.com                         |
|     TERMSRV/UFC-DC1                                                                            |
|     TERMSRV/rt-dc.rt.securecompany.com                                                           |
|     DNS/rt-dc.rt.securecompany.com                                                               |
|     GC/rt-dc.rt.securecompany.com/funcorp.local                                                  |
|     RestrictedKrbHost/rt-dc.rt.securecompany.com                                                 |
|     RestrictedKrbHost/UFC-DC1                                                                  |
|     RPC/5873d6ca-8398-42dc-8c56-885c4ed4cd2b._msdcs.funcorp.local                              |
|     HOST/UFC-DC1/USFUN                                                                         |
|     HOST/rt-dc.rt.securecompany.com/USFUN                                                        |
|     HOST/UFC-DC1                                                                               |
|     HOST/rt-dc.rt.securecompany.com                                                              |
|     HOST/rt-dc.rt.securecompany.com/rt.securecompany.com                                             |
|     E3514235-4B06-11D1-AB04-00C04FC2DCD2/5873d6ca-8398-42dc-8c56-885c4ed4cd2b/rt.securecompany.com |
|     ldap/UFC-DC1/USFUN                                                                         |
|     ldap/5873d6ca-8398-42dc-8c56-885c4ed4cd2b._msdcs.funcorp.local                             |
|     ldap/rt-dc.rt.securecompany.com/USFUN                                                        |
|     ldap/UFC-DC1                                                                               |
|     ldap/rt-dc.rt.securecompany.com                                                              |
|     ldap/rt-dc.rt.securecompany.com/rt.securecompany.com                                             |
| Operating System : Windows Server 2016 Standard                                                |
| Operating System Version : 10.0 (14393)                                                        |
| UAC Flag : SERVER_TRUST_ACCOUNT,DONT_EXPIRE_PASSWD,TRUSTED_FOR_DELEGATION                      |
| User Account Control : 598016                                                                  |
| Object GUID : 68412feb-0296-4eeb-8886-ae49b2779f64                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1000                                    |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20190201062300.0Z                                                               |
| When Changed : 20220419133214.0Z                                                               |
| Password Last Set : 2019-05-16 08:10:20.0355455 +0000 UTC                                      |
| Last Log On : 2022-04-27 06:36:05.8854294 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.0646402 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0012293 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
```

To query the ```Trust Attributes``` of the current domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command domain-trust 
+------------------------------------------------------------------------------------+
| TRUSTED DOMAIN                                                                     |
+------------------------------------------------------------------------------------+
| DN : CN=funcorp.local,CN=System,DC=us,DC=funcorp,DC=local                          |
| CN : funcorp.local                                                                 |
| Instance Type : 4                                                                  |
| Object Class : top, leaf, trustedDomain                                            |
| Object Category : CN=Trusted-Domain,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Trust Type : UPLEVEL                                                               |
| Trust Direction : INBOUND,OUTBOUND                                                 |
| Trust Attribute : WITHIN_FOREST                                                    |
| Object GUID : 031b05c5-4d15-49dc-81a5-92020926921d                                 |
| When Created : 20190201062154.0Z                                                   |
| When Changed : 20220419133214.0Z                                                   |
| uSNChanged : 1601-01-01 00:00:00.0646406 +0000 UTC                                 |
| uSNCreated : 1601-01-01 00:00:00.0008242 +0000 UTC                                 |
+------------------------------------------------------------------------------------+
```

To list all ```Users``` in the domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command users 
+-------------------------+
| USERS                   |
+-------------------------+
|  - Administrator        |
|  - Guest                |
|  - DefaultAccount       |
|  - UFC-DC1$             |
|  - krbtgt               |
|  - FUNCORP$             |
|  - UFC-WEBPROD$         |
|  - UFC-DBPROD$          |
|  - UFC-SQLDEV$          |
|  - UFC-APP1$            |
|  - UFC-DB1$             |
|  - UFC-JUMPSRV$         |
|  - appadmin             |
|  - dbadmin              |
|  - dbservice            |
|  - db1user              |
|  - servicesadmin        |
|  - sqldevadmin          |
|  - sqlreportuser        |
|  - jumpsrvadmin         |
|  - webprodadmin         |
|  - dbprodadmin          |
|  - Woming               |
|  - Andrescrove          |
|  - Onnithashe           |
|  - Whirosed             |
|  - Addren               |

...

|  - PA-USER117$          |
|  - PA-USER118$          |
|  - PA-USER119$          |
|  - pastudent120         |
|  - pastudent121         |
|  - PA-USER120$          |
|  - PA-USER121$          |
+-------------------------+
| 2011 Domain Users Found |
+-------------------------+
```

To list all ```Users``` with detailed properties in the domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command users -filter full-data >out.txt   
+------------------------------------------------------------------------------------------------+
| USERS                                                                                          |
+------------------------------------------------------------------------------------------------+
| DN : CN=Administrator,CN=Users,DC=us,DC=funcorp,DC=local                                       |
| SAM Account Name : Administrator                                                               |
| SAM Account Type : 805306368                                                                   |
| CN : Administrator                                                                             |
| Description : Built-in account for administering the computer/domain                           |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user                                         |
| Object Category : CN=Person,CN=Schema,CN=Configuration,DC=funcorp,DC=local                     |
| Member Of :                                                                                    |
|     CN=Group Policy Creator Owners,CN=Users,DC=us,DC=funcorp,DC=local                          |
|     CN=Domain Admins,CN=Users,DC=us,DC=funcorp,DC=local                                        |
|     CN=Administrators,CN=Builtin,DC=us,DC=funcorp,DC=local                                     |
| UAC Flag : NORMAL_ACCOUNT,DONT_EXPIRE_PASSWD                                                   |
| User Account Control : 66048                                                                   |
| Object GUID : b074a3fe-ff69-498f-920e-a726fb6d696c                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-500                                     |
| Bad Password Count : 2                                                                         |
| Bad Password Time : 2022-04-27 13:06:22.1553123 +0000 UTC                                      |
| When Created : 20190201062154.0Z                                                               |
| When Changed : 20220419133235.0Z                                                               |
| Password Last Set : 2019-03-28 09:57:28.9673316 +0000 UTC                                      |
| Last Log On : 2022-04-19 13:32:36.0307154 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 1601-01-01 00:00:00 +0000 UTC                                             |
| uSNChanged : 1601-01-01 00:00:00.0646443 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0008196 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN=Guest,CN=Users,DC=us,DC=funcorp,DC=local                                               |
| SAM Account Name : Guest                                                                       |
| SAM Account Type : 805306368                                                                   |
| CN : Guest                                                                                     |
| Description : Built-in account for guest access to the computer/domain                         |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user                                         |
| Object Category : CN=Person,CN=Schema,CN=Configuration,DC=funcorp,DC=local                     |
| Member Of :                                                                                    |
|     CN=Guests,CN=Builtin,DC=us,DC=funcorp,DC=local                                             |
| UAC Flag : ACCOUNT_DISABLED,PASSWD_NOTREQD,NORMAL_ACCOUNT,DONT_EXPIRE_PASSWD                   |
| User Account Control : 66082                                                                   |
| Object GUID : 1d499bea-a5e2-483e-b056-b4ca9199c239                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-501                                     |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20190201062154.0Z                                                               |
| When Changed : 20190201062154.0Z                                                               |
| Password Last Set : 1601-01-01 00:00:00 +0000 UTC                                              |
| Last Log On : 1601-01-01 00:00:00 +0000 UTC                                                    |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.0008197 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0008197 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN=DefaultAccount,CN=Users,DC=us,DC=funcorp,DC=local                                      |
| SAM Account Name : DefaultAccount                                                              |
| SAM Account Type : 805306368                                                                   |
| CN : DefaultAccount                                                                            |
| Description : A user account managed by the system.                                            |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user                                         |
| Object Category : CN=Person,CN=Schema,CN=Configuration,DC=funcorp,DC=local                     |
| Member Of :                                                                                    |
|     CN=System Managed Accounts Group,CN=Builtin,DC=us,DC=funcorp,DC=local                      |
| UAC Flag : ACCOUNT_DISABLED,PASSWD_NOTREQD,NORMAL_ACCOUNT,DONT_EXPIRE_PASSWD                   |
| User Account Control : 66082                                                                   |
| Object GUID : b6007ba5-91e8-44fd-8f23-78cbf226f185                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-503                                     |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20190201062154.0Z                                                               |
| When Changed : 20190201062154.0Z                                                               |
| Password Last Set : 1601-01-01 00:00:00 +0000 UTC                                              |
| Last Log On : 1601-01-01 00:00:00 +0000 UTC                                                    |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.0008198 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0008198 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+

...

+------------------------------------------------------------------------------------------------+
| DN : CN=PA-USER120,CN=Computers,DC=us,DC=funcorp,DC=local                                      |
| SAM Account Name : PA-USER120$                                                                 |
| SAM Account Type : 805306369                                                                   |
| CN : PA-USER120                                                                                |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user, computer                               |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local                   |
| DNS Host Name : PA-User120.rt.securecompany.com                                                    |
| Service Principal Name (SPN) :                                                                 |
|     WSMAN/PA-User120                                                                           |
|     WSMAN/PA-User120.rt.securecompany.com                                                          |
|     TERMSRV/PA-USER120                                                                         |
|     TERMSRV/PA-User120.rt.securecompany.com                                                        |
|     RestrictedKrbHost/PA-USER120                                                               |
|     HOST/PA-USER120                                                                            |
|     RestrictedKrbHost/PA-User120.rt.securecompany.com                                              |
|     HOST/PA-User120.rt.securecompany.com                                                           |
| Operating System : Windows Server 2016 Standard                                                |
| Operating System Version : 10.0 (14393)                                                        |
| UAC Flag : WORKSTATION_ACCOUNT                                                                 |
| User Account Control : 4096                                                                    |
| Object GUID : 8876bdf1-8add-4a6a-b7bd-1e553442ff65                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-38103                                   |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20210209164458.0Z                                                               |
| When Changed : 20211222074651.0Z                                                               |
| Password Last Set : 2021-02-09 16:44:58.8997332 +0000 UTC                                      |
| Last Log On : 2021-12-22 07:50:03.0833686 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.0646155 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0645944 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN=PA-USER121,CN=Computers,DC=us,DC=funcorp,DC=local                                      |
| SAM Account Name : PA-USER121$                                                                 |
| SAM Account Type : 805306369                                                                   |
| CN : PA-USER121                                                                                |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user, computer                               |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local                   |
| DNS Host Name : PA-User121.rt.securecompany.com                                                    |
| Service Principal Name (SPN) :                                                                 |
|     WSMAN/PA-User121                                                                           |
|     WSMAN/PA-User121.rt.securecompany.com                                                          |
|     TERMSRV/PA-USER121                                                                         |
|     TERMSRV/PA-User121.rt.securecompany.com                                                        |
|     RestrictedKrbHost/PA-USER121                                                               |
|     HOST/PA-USER121                                                                            |
|     RestrictedKrbHost/PA-User121.rt.securecompany.com                                              |
|     HOST/PA-User121.rt.securecompany.com                                                           |
| Operating System : Windows Server 2016 Standard                                                |
| Operating System Version : 10.0 (14393)                                                        |
| UAC Flag : WORKSTATION_ACCOUNT                                                                 |
| User Account Control : 4096                                                                    |
| Object GUID : 41fdacff-fe2d-405a-ad16-6afef2d79739                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-38104                                   |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20210209164511.0Z                                                               |
| When Changed : 20211222074716.0Z                                                               |
| Password Last Set : 2021-02-09 16:45:11.2259445 +0000 UTC                                      |
| Last Log On : 2021-12-22 07:50:28.9690176 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.064618 +0000 UTC                                              |
| uSNCreated : 1601-01-01 00:00:00.0645959 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
```

To list all attributes of ```specific user```,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command users -name servicesadmin
+----------------------------------------------------------------------------+
| USERS                                                                      |
+----------------------------------------------------------------------------+
| DN : CN=services admin,CN=Users,DC=us,DC=funcorp,DC=local                  |
| SAM Account Name : servicesadmin                                           |
| SAM Account Type : 805306368                                               |
| CN : services admin                                                        |
| Given Name : services                                                      |
| SN : admin                                                                 |
| Description : Account to be used for services which need high privileges.  |
| Instance Type : 4                                                          |
| Object Class : top, person, organizationalPerson, user                     |
| Object Category : CN=Person,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Member Of :                                                                |
|     CN=Domain Admins,CN=Users,DC=us,DC=funcorp,DC=local                    |
| UAC Flag : NORMAL_ACCOUNT,DONT_EXPIRE_PASSWD,NOT_DELEGATED                 |
| User Account Control : 1114624                                             |
| Object GUID : 4c243de5-495f-4367-a82f-341916c85cd2                         |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1122                |
| Bad Password Count : 0                                                     |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                          |
| When Created : 20190204063650.0Z                                           |
| When Changed : 20190204064328.0Z                                           |
| Password Last Set : 2019-02-04 06:36:50.1431059 +0000 UTC                  |
| Last Log On : 1601-01-01 00:00:00 +0000 UTC                                |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                               |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                |
| uSNChanged : 1601-01-01 00:00:00.0066717 +0000 UTC                         |
| uSNCreated : 1601-01-01 00:00:00.0066567 +0000 UTC                         |
+----------------------------------------------------------------------------+
```

To list ```membership``` of ```specific user```,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command users -name servicesadmin -filter membership
+---------------------------------------------------------+
| USERS                                                   |
+---------------------------------------------------------+
| Member Of :                                             |
|     CN=Domain Admins,CN=Users,DC=us,DC=funcorp,DC=local |
+---------------------------------------------------------+
```

To list all available ```Computers``` in this domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command computers
+---------------------------+
| COMPUTERS                 |
+---------------------------+
|  - UFC-DC1$               |
|  - UFC-WEBPROD$           |
|  - UFC-DBPROD$            |
|  - UFC-SQLDEV$            |
|  - UFC-APP1$              |
|  - UFC-DB1$               |
|  - UFC-JUMPSRV$           |
|  - PA-ADMIN$              |
|  - PA-USER112$            |
|  - PA-USER113$            |
|  - PA-USER114$            |
|  - PA-USER115$            |
|  - PA-USER116$            |
|  - PA-USER117$            |
|  - PA-USER118$            |
|  - PA-USER119$            |
|  - PA-USER120$            |
|  - PA-USER121$            |
+---------------------------+
| 18 Domain Computers Found |
+---------------------------+
```

To list all computers in domain with attributes,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command computers -filter full-data >out.txt
+------------------------------------------------------------------------------------------------+
| COMPUTERS                                                                                      |
+------------------------------------------------------------------------------------------------+
| DN : CN=UFC-DC1,OU=Domain Controllers,DC=us,DC=funcorp,DC=local                                |
| SAM Account Name : UFC-DC1$                                                                    |
| SAM Account Type : 805306369                                                                   |
| CN : UFC-DC1                                                                                   |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user, computer                               |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local                   |
| DNS Host Name : rt-dc.rt.securecompany.com                                                       |
| Service Principal Name (SPN) :                                                                 |
|     ldap/rt-dc.rt.securecompany.com/DomainDnsZones.rt.securecompany.com                              |
|     ldap/rt-dc.rt.securecompany.com/ForestDnsZones.funcorp.local                                 |
|     Dfsr-12F9A27C-BF97-4787-9364-D31B6C55EB04/rt-dc.rt.securecompany.com                         |
|     TERMSRV/UFC-DC1                                                                            |
|     TERMSRV/rt-dc.rt.securecompany.com                                                           |
|     DNS/rt-dc.rt.securecompany.com                                                               |
|     GC/rt-dc.rt.securecompany.com/funcorp.local                                                  |
|     RestrictedKrbHost/rt-dc.rt.securecompany.com                                                 |
|     RestrictedKrbHost/UFC-DC1                                                                  |
|     RPC/5873d6ca-8398-42dc-8c56-885c4ed4cd2b._msdcs.funcorp.local                              |
|     HOST/UFC-DC1/USFUN                                                                         |
|     HOST/rt-dc.rt.securecompany.com/USFUN                                                        |
|     HOST/UFC-DC1                                                                               |
|     HOST/rt-dc.rt.securecompany.com                                                              |
|     HOST/rt-dc.rt.securecompany.com/rt.securecompany.com                                             |
|     E3514235-4B06-11D1-AB04-00C04FC2DCD2/5873d6ca-8398-42dc-8c56-885c4ed4cd2b/rt.securecompany.com |
|     ldap/UFC-DC1/USFUN                                                                         |
|     ldap/5873d6ca-8398-42dc-8c56-885c4ed4cd2b._msdcs.funcorp.local                             |
|     ldap/rt-dc.rt.securecompany.com/USFUN                                                        |
|     ldap/UFC-DC1                                                                               |
|     ldap/rt-dc.rt.securecompany.com                                                              |
|     ldap/rt-dc.rt.securecompany.com/rt.securecompany.com                                             |
| Operating System : Windows Server 2016 Standard                                                |
| Operating System Version : 10.0 (14393)                                                        |
| UAC Flag : SERVER_TRUST_ACCOUNT,DONT_EXPIRE_PASSWD,TRUSTED_FOR_DELEGATION                      |
| User Account Control : 598016                                                                  |
| Object GUID : 68412feb-0296-4eeb-8886-ae49b2779f64                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1000                                    |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20190201062300.0Z                                                               |
| When Changed : 20220419133214.0Z                                                               |
| Password Last Set : 2019-05-16 08:10:20.0355455 +0000 UTC                                      |
| Last Log On : 2022-04-27 14:36:06.0316569 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.0646402 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0012293 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+

...

+------------------------------------------------------------------------------------------------+
| DN : CN=PA-USER120,CN=Computers,DC=us,DC=funcorp,DC=local                                      |
| SAM Account Name : PA-USER120$                                                                 |
| SAM Account Type : 805306369                                                                   |
| CN : PA-USER120                                                                                |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user, computer                               |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local                   |
| DNS Host Name : PA-User120.rt.securecompany.com                                                    |
| Service Principal Name (SPN) :                                                                 |
|     WSMAN/PA-User120                                                                           |
|     WSMAN/PA-User120.rt.securecompany.com                                                          |
|     TERMSRV/PA-USER120                                                                         |
|     TERMSRV/PA-User120.rt.securecompany.com                                                        |
|     RestrictedKrbHost/PA-USER120                                                               |
|     HOST/PA-USER120                                                                            |
|     RestrictedKrbHost/PA-User120.rt.securecompany.com                                              |
|     HOST/PA-User120.rt.securecompany.com                                                           |
| Operating System : Windows Server 2016 Standard                                                |
| Operating System Version : 10.0 (14393)                                                        |
| UAC Flag : WORKSTATION_ACCOUNT                                                                 |
| User Account Control : 4096                                                                    |
| Object GUID : 8876bdf1-8add-4a6a-b7bd-1e553442ff65                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-38103                                   |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20210209164458.0Z                                                               |
| When Changed : 20211222074651.0Z                                                               |
| Password Last Set : 2021-02-09 16:44:58.8997332 +0000 UTC                                      |
| Last Log On : 2021-12-22 07:50:03.0833686 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.0646155 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0645944 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN=PA-USER121,CN=Computers,DC=us,DC=funcorp,DC=local                                      |
| SAM Account Name : PA-USER121$                                                                 |
| SAM Account Type : 805306369                                                                   |
| CN : PA-USER121                                                                                |
| Instance Type : 4                                                                              |
| Object Class : top, person, organizationalPerson, user, computer                               |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local                   |
| DNS Host Name : PA-User121.rt.securecompany.com                                                    |
| Service Principal Name (SPN) :                                                                 |
|     WSMAN/PA-User121                                                                           |
|     WSMAN/PA-User121.rt.securecompany.com                                                          |
|     TERMSRV/PA-USER121                                                                         |
|     TERMSRV/PA-User121.rt.securecompany.com                                                        |
|     RestrictedKrbHost/PA-USER121                                                               |
|     HOST/PA-USER121                                                                            |
|     RestrictedKrbHost/PA-User121.rt.securecompany.com                                              |
|     HOST/PA-User121.rt.securecompany.com                                                           |
| Operating System : Windows Server 2016 Standard                                                |
| Operating System Version : 10.0 (14393)                                                        |
| UAC Flag : WORKSTATION_ACCOUNT                                                                 |
| User Account Control : 4096                                                                    |
| Object GUID : 41fdacff-fe2d-405a-ad16-6afef2d79739                                             |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-38104                                   |
| Bad Password Count : 0                                                                         |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                                              |
| When Created : 20210209164511.0Z                                                               |
| When Changed : 20211222074716.0Z                                                               |
| Password Last Set : 2021-02-09 16:45:11.2259445 +0000 UTC                                      |
| Last Log On : 2021-12-22 07:50:28.9690176 +0000 UTC                                            |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                                   |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                                    |
| uSNChanged : 1601-01-01 00:00:00.064618 +0000 UTC                                              |
| uSNCreated : 1601-01-01 00:00:00.0645959 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
```

To list all attributes of specific computer,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command computers -name ufc-sqldev
+------------------------------------------------------------------------------+
| COMPUTERS                                                                    |
+------------------------------------------------------------------------------+
| DN : CN=UFC-SQLDEV,OU=Servers,DC=us,DC=funcorp,DC=local                      |
| SAM Account Name : UFC-SQLDEV$                                               |
| SAM Account Type : 805306369                                                 |
| CN : UFC-SQLDEV                                                              |
| Instance Type : 4                                                            |
| Object Class : top, person, organizationalPerson, user, computer             |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| DNS Host Name : UFC-SQLDev.rt.securecompany.com                                  |
| Service Principal Name (SPN) :                                               |
|     MSSQLSvc/UFC-SQLDev.rt.securecompany.com:1433                                |
|     MSSQLSvc/UFC-SQLDev.rt.securecompany.com                                     |
|     WSMAN/UFC-SQLDev                                                         |
|     WSMAN/UFC-SQLDev.rt.securecompany.com                                        |
|     TERMSRV/UFC-SQLDEV                                                       |
|     TERMSRV/UFC-SQLDev.rt.securecompany.com                                      |
|     RestrictedKrbHost/UFC-SQLDEV                                             |
|     HOST/UFC-SQLDEV                                                          |
|     RestrictedKrbHost/UFC-SQLDev.rt.securecompany.com                            |
|     HOST/UFC-SQLDev.rt.securecompany.com                                         |
| Operating System : Windows Server 2016 Standard                              |
| Operating System Version : 10.0 (14393)                                      |
| UAC Flag : WORKSTATION_ACCOUNT,DONT_EXPIRE_PASSWD                            |
| User Account Control : 69632                                                 |
| Object GUID : 19f7e5f0-3753-43ae-87a3-a19fc7e92247                           |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1106                  |
| Bad Password Count : 0                                                       |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                            |
| When Created : 20190201071830.0Z                                             |
| When Changed : 20220419133235.0Z                                             |
| Password Last Set : 2019-06-17 11:00:32.1796335 +0000 UTC                    |
| Last Log On : 2022-04-27 14:54:38.7180042 +0000 UTC                          |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                 |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                  |
| uSNChanged : 1601-01-01 00:00:00.0646442 +0000 UTC                           |
| uSNCreated : 1601-01-01 00:00:00.0013204 +0000 UTC                           |
+------------------------------------------------------------------------------+
```

To list all group available in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command groups
+--------------------------------------------+
| GROUPS                                     |
+--------------------------------------------+
|  - Administrators                          |
|  - Users                                   |
|  - Guests                                  |
|  - Print Operators                         |
|  - Backup Operators                        |
|  - Replicator                              |
|  - Remote Desktop Users                    |
|  - Network Configuration Operators         |
|  - Performance Monitor Users               |
|  - Performance Log Users                   |
|  - Distributed COM Users                   |
|  - IIS_IUSRS                               |
|  - Cryptographic Operators                 |
|  - Event Log Readers                       |
|  - Certificate Service DCOM Access         |
|  - RDS Remote Access Servers               |
|  - RDS Endpoint Servers                    |
|  - RDS Management Servers                  |
|  - Hyper-V Administrators                  |
|  - Access Control Assistance Operators     |
|  - Remote Management Users                 |
|  - System Managed Accounts Group           |
|  - Storage Replica Administrators          |
|  - Domain Computers                        |
|  - Domain Controllers                      |
|  - Cert Publishers                         |
|  - Domain Admins                           |
|  - Domain Users                            |
|  - Domain Guests                           |
|  - Group Policy Creator Owners             |
|  - RAS and IAS Servers                     |
|  - Server Operators                        |
|  - Account Operators                       |
|  - Pre-Windows 2000 Compatible Access      |
|  - Windows Authorization Access Group      |
|  - Terminal Server License Servers         |
|  - Allowed RODC Password Replication Group |
|  - Denied RODC Password Replication Group  |
|  - Read-only Domain Controllers            |
|  - Cloneable Domain Controllers            |
|  - Protected Users                         |
|  - Key Admins                              |
|  - DnsAdmins                               |
|  - DnsUpdateProxy                          |
|  - RDPUsers                                |
|  - ProductionManagers                      |
+--------------------------------------------+
| 46 Domain Groups Found                     |
+--------------------------------------------+
```

To list all properties of groups in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command groups -filter full-data 
+------------------------------------------------------------------------------------------------------+
| GROUPS                                                                                               |
+------------------------------------------------------------------------------------------------------+
| DN : CN=Administrators,CN=Builtin,DC=us,DC=funcorp,DC=local                                          |
| SAM Account Name : Administrators                                                                    |
| SAM Account Type : 536870912                                                                         |
| CN : Administrators                                                                                  |
| Description : Administrators have complete and unrestricted access to the computer/domain            |
| Instance Type : 4                                                                                    |
| Object Class : top, group                                                                            |
| Object Category : CN=Group,CN=Schema,CN=Configuration,DC=funcorp,DC=local                            |
| Object GUID : ec0a24b8-33e1-4d28-ba8e-1d32a8126ee8                                                   |
| Object SID : S-1-5-32-544                                                                            |
| When Created : 20190201062154.0Z                                                                     |
| When Changed : 20190201065309.0Z                                                                     |
| uSNChanged : 1601-01-01 00:00:00.001301 +0000 UTC                                                    |
| uSNCreated : 1601-01-01 00:00:00.00082 +0000 UTC                                                     |
+------------------------------------------------------------------------------------------------------+
| DN : CN=Users,CN=Builtin,DC=us,DC=funcorp,DC=local                                                   |
| SAM Account Name : Users                                                                             |
| SAM Account Type : 536870912                                                                         |
| CN : Users                                                                                           |
| Description : Users are prevented from making accidental or intentional system-wide changes and can  |
| run most applications                                                                                |
| Instance Type : 4                                                                                    |
| Object Class : top, group                                                                            |
| Object Category : CN=Group,CN=Schema,CN=Configuration,DC=funcorp,DC=local                            |
| Object GUID : 531b56ca-7219-41e6-876b-9884c8b9fe46                                                   |
| Object SID : S-1-5-32-545                                                                            |
| When Created : 20190201062154.0Z                                                                     |
| When Changed : 20190201062300.0Z                                                                     |
| uSNChanged : 1601-01-01 00:00:00.0012348 +0000 UTC                                                   |
| uSNCreated : 1601-01-01 00:00:00.0008203 +0000 UTC                                                   |
+------------------------------------------------------------------------------------------------------+
| DN : CN=Guests,CN=Builtin,DC=us,DC=funcorp,DC=local                                                  |
| SAM Account Name : Guests                                                                            |
| SAM Account Type : 536870912                                                                         |
| CN : Guests                                                                                          |
| Description : Guests have the same access as members of the Users group by default, except for the G |
| uest account which is further restricted                                                             |
| Instance Type : 4                                                                                    |
| Object Class : top, group                                                                            |
| Object Category : CN=Group,CN=Schema,CN=Configuration,DC=funcorp,DC=local                            |
| Object GUID : 80be0504-53e7-4aba-bb6e-66096fd007c4                                                   |
| Object SID : S-1-5-32-546                                                                            |
| When Created : 20190201062154.0Z                                                                     |
| When Changed : 20190201062300.0Z                                                                     |
| uSNChanged : 1601-01-01 00:00:00.001235 +0000 UTC                                                    |
| uSNCreated : 1601-01-01 00:00:00.0008209 +0000 UTC                                                   |
+------------------------------------------------------------------------------------------------------+

...

+------------------------------------------------------------------------------------------------------+
| DN : CN=RDP Users,CN=Users,DC=us,DC=funcorp,DC=local                                                 |
| SAM Account Name : RDPUsers                                                                          |
| SAM Account Type : 268435456                                                                         |
| CN : RDP Users                                                                                       |
| Description : RDP Users Group                                                                        |
| Instance Type : 4                                                                                    |
| Object Class : top, group                                                                            |
| Object Category : CN=Group,CN=Schema,CN=Configuration,DC=funcorp,DC=local                            |
| Object GUID : 34affcf8-81eb-434b-b7a0-05da0c2837e9                                                   |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1123                                          |
| When Created : 20190204063655.0Z                                                                     |
| When Changed : 20210209164337.0Z                                                                     |
| uSNChanged : 1601-01-01 00:00:00.0645918 +0000 UTC                                                   |
| uSNCreated : 1601-01-01 00:00:00.006658 +0000 UTC                                                    |
+------------------------------------------------------------------------------------------------------+
| DN : CN=ProductionManagers,CN=Users,DC=us,DC=funcorp,DC=local                                        |
| SAM Account Name : ProductionManagers                                                                |
| SAM Account Type : 268435456                                                                         |
| CN : ProductionManagers                                                                              |
| Description : Production Managers Group                                                              |
| Instance Type : 4                                                                                    |
| Object Class : top, group                                                                            |
| Object Category : CN=Group,CN=Schema,CN=Configuration,DC=funcorp,DC=local                            |
| Object GUID : 0fdade69-d5d8-410a-81ab-dae18fe390fe                                                   |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1127                                          |
| When Created : 20190204064307.0Z                                                                     |
| When Changed : 20200921091448.0Z                                                                     |
| uSNChanged : 1601-01-01 00:00:00.0568153 +0000 UTC                                                   |
| uSNCreated : 1601-01-01 00:00:00.0066699 +0000 UTC                                                   |
+------------------------------------------------------------------------------------------------------+
```

To list all attributes of a specific group in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command groups -name "ProductionManagers"
| GROUPS                                                                    |
+---------------------------------------------------------------------------+
| DN : CN=ProductionManagers,CN=Users,DC=us,DC=funcorp,DC=local             |
| SAM Account Name : ProductionManagers                                     |
| SAM Account Type : 268435456                                              |
| CN : ProductionManagers                                                   |
| Description : Production Managers Group                                   |
| Instance Type : 4                                                         |
| Object Class : top, group                                                 |
| Object Category : CN=Group,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 0fdade69-d5d8-410a-81ab-dae18fe390fe                        |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1127               |
| When Created : 20190204064307.0Z                                          |
| When Changed : 20200921091448.0Z                                          |
| uSNChanged : 1601-01-01 00:00:00.0568153 +0000 UTC                        |
| uSNCreated : 1601-01-01 00:00:00.0066699 +0000 UTC                        |
+---------------------------------------------------------------------------+
```

To list all members of a group in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command groups -name "Domain Admins" -filter membership
 +--------------------------+
| GROUPS                   |
+--------------------------+
| Domain users in group :  |
|  - Administrator         |
|  - servicesadmin         |
+--------------------------+
```

To list all users in domain who never logged on

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command never-loggedon 
+----------------------+
| USERS NEVER LOGGEDON |
+----------------------+
|  - Guest             |
|  - DefaultAccount    |
|  - krbtgt            |
|  - FUNCORP$          |
|  - db1user           |
|  - servicesadmin     |
|  - sqlreportuser     |
|  - Woming            |
|  - Andrescrove       |
|  - Onnithashe        |

...

|  - Thempern1982      |
|  - Voymber           |
|  - Thintich          |
|  - Stolven1987       |
|  - Weirche           |
|  - Porticed          |
|  - Ling1987          |
+----------------------+
```

To list all GPOs in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command gpo
+------------------------------------------------------------------------------------------------+
| GROUP POLICY OBJECTS                                                                           |
+------------------------------------------------------------------------------------------------+
| DN : CN={31B2F340-016D-11D2-945F-00C04FB984F9},CN=Policies,CN=System,DC=us,DC=funcorp,DC=local |
| CN : {31B2F340-016D-11D2-945F-00C04FB984F9}                                                    |
| Instance Type : 4                                                                              |
| Object Class : top, container, groupPolicyContainer                                            |
| Object Category : CN=Group-Policy-Container,CN=Schema,CN=Configuration,DC=funcorp,DC=local     |
| Object GUID : 012ce41a-8f84-4526-9a8c-7d04bbfe597b                                             |
| When Created : 20190201062151.0Z                                                               |
| When Changed : 20190523051909.0Z                                                               |
| uSNChanged : 1601-01-01 00:00:00.0345563 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0008016 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN={6AC1786C-016F-11D2-945F-00C04fB984F9},CN=Policies,CN=System,DC=us,DC=funcorp,DC=local |
| CN : {6AC1786C-016F-11D2-945F-00C04fB984F9}                                                    |
| Instance Type : 4                                                                              |
| Object Class : top, container, groupPolicyContainer                                            |
| Object Category : CN=Group-Policy-Container,CN=Schema,CN=Configuration,DC=funcorp,DC=local     |
| Object GUID : 246c7b3a-98d7-467f-a459-0f08ac6e2d5a                                             |
| When Created : 20190201062151.0Z                                                               |
| When Changed : 20190523052819.0Z                                                               |
| uSNChanged : 1601-01-01 00:00:00.0345896 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0008019 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN={603ABE02-C554-49B1-A462-2FF89BC61CB2},CN=Policies,CN=System,DC=us,DC=funcorp,DC=local |
| CN : {603ABE02-C554-49B1-A462-2FF89BC61CB2}                                                    |
| Instance Type : 4                                                                              |
| Object Class : top, container, groupPolicyContainer                                            |
| Object Category : CN=Group-Policy-Container,CN=Schema,CN=Configuration,DC=funcorp,DC=local     |
| Object GUID : f3249764-3bf1-4f92-8150-daae205b51a6                                             |
| When Created : 20190206063611.0Z                                                               |
| When Changed : 20190523052328.0Z                                                               |
| uSNChanged : 1601-01-01 00:00:00.0345727 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0105756 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
| DN : CN={B822494A-DD6A-4E96-A2BB-944E397208A1},CN=Policies,CN=System,DC=us,DC=funcorp,DC=local |
| CN : {B822494A-DD6A-4E96-A2BB-944E397208A1}                                                    |
| Instance Type : 4                                                                              |
| Object Class : top, container, groupPolicyContainer                                            |
| Object Category : CN=Group-Policy-Container,CN=Schema,CN=Configuration,DC=funcorp,DC=local     |
| Object GUID : f6ca2212-eb6f-4eda-9ee8-26176f85ec5e                                             |
| When Created : 20190206123933.0Z                                                               |
| When Changed : 20190523052930.0Z                                                               |
| uSNChanged : 1601-01-01 00:00:00.0345931 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.012257 +0000 UTC                                              |
+------------------------------------------------------------------------------------------------+
| DN : CN={C95D8D85-BFE6-453E-9668-E31379106EB0},CN=Policies,CN=System,DC=us,DC=funcorp,DC=local |
| CN : {C95D8D85-BFE6-453E-9668-E31379106EB0}                                                    |
| Instance Type : 4                                                                              |
| Object Class : top, container, groupPolicyContainer                                            |
| Object Category : CN=Group-Policy-Container,CN=Schema,CN=Configuration,DC=funcorp,DC=local     |
| Object GUID : 81460d95-ea9e-4dc0-afe9-feb0b7652ab9                                             |
| When Created : 20190206124702.0Z                                                               |
| When Changed : 20190523052412.0Z                                                               |
| uSNChanged : 1601-01-01 00:00:00.0345757 +0000 UTC                                             |
| uSNCreated : 1601-01-01 00:00:00.0122698 +0000 UTC                                             |
+------------------------------------------------------------------------------------------------+
```

To list all OUs in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command ou
+-----------------------------------------------------------------------------------------+
| ORGANIZATIONAL UNITS                                                                    |
+-----------------------------------------------------------------------------------------+
| DN : OU=Domain Controllers,DC=us,DC=funcorp,DC=local                                    |
| OU : Domain Controllers                                                                 |
| Description : Default container for domain controllers                                  |
| Instance Type : 4                                                                       |
| Object Class : top, organizationalUnit                                                  |
| Object Category : CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 20d12650-d90e-4c09-ae48-541b34075b33                                      |
| When Created : 20190201062152.0Z                                                        |
| When Changed : 20190201062152.0Z                                                        |
| uSNChanged : 1601-01-01 00:00:00.0008147 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0008147 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
| DN : OU=StudentMachines,DC=us,DC=funcorp,DC=local                                       |
| OU : StudentMachines                                                                    |
| Instance Type : 4                                                                       |
| Object Class : top, organizationalUnit                                                  |
| Object Category : CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 6f17609c-13e3-48ac-9af5-ec93af18791c                                      |
| When Created : 20190204063650.0Z                                                        |
| When Changed : 20190206123934.0Z                                                        |
| uSNChanged : 1601-01-01 00:00:00.0122576 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0066575 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
| DN : OU=Applocked,DC=us,DC=funcorp,DC=local                                             |
| OU : Applocked                                                                          |
| Instance Type : 4                                                                       |
| Object Class : top, organizationalUnit                                                  |
| Object Category : CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 99e5e32d-2f78-48c5-8430-dd4a15f5c272                                      |
| When Created : 20190204064247.0Z                                                        |
| When Changed : 20190206063611.0Z                                                        |
| uSNChanged : 1601-01-01 00:00:00.0105762 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0066692 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
| DN : OU=ActiveUsers,DC=us,DC=funcorp,DC=local                                           |
| OU : ActiveUsers                                                                        |
| Instance Type : 4                                                                       |
| Object Class : top, organizationalUnit                                                  |
| Object Category : CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 6ee45b8d-0564-4277-b455-2c59ffdb1652                                      |
| When Created : 20190206123232.0Z                                                        |
| When Changed : 20190206123325.0Z                                                        |
| uSNChanged : 1601-01-01 00:00:00.0110676 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0110661 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
| DN : OU=US,OU=ActiveUsers,DC=us,DC=funcorp,DC=local                                     |
| OU : US                                                                                 |
| Instance Type : 4                                                                       |
| Object Class : top, organizationalUnit                                                  |
| Object Category : CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 23945556-cb8d-42f6-a911-c517c9576dc7                                      |
| When Created : 20190206123325.0Z                                                        |
| When Changed : 20190206123325.0Z                                                        |
| uSNChanged : 1601-01-01 00:00:00.0110677 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0110675 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
| DN : OU=Servers,DC=us,DC=funcorp,DC=local                                               |
| OU : Servers                                                                            |
| Instance Type : 4                                                                       |
| Object Class : top, organizationalUnit                                                  |
| Object Category : CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| Object GUID : 12a1190c-e24e-4ad4-b084-92d9e47e58ed                                      |
| When Created : 20190206124631.0Z                                                        |
| When Changed : 20190206124702.0Z                                                        |
| uSNChanged : 1601-01-01 00:00:00.0122704 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0122682 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
```

To list all ```MS-SQL Servers``` in domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command ms-sql 
+----------------------+
| MS-SQL SERVERS       |
+----------------------+
|  - UFC-DBPROD$       |
|  - UFC-SQLDEV$       |
+----------------------+
```

To list all MS-SQL Servers with all attributes,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command ms-sql -filter full-data 
+------------------------------------------------------------------------------+
| MS-SQL SERVERS                                                               |
+------------------------------------------------------------------------------+
| DN : CN=UFC-DBPROD,OU=Servers,DC=us,DC=funcorp,DC=local                      |
| SAM Account Name : UFC-DBPROD$                                               |
| SAM Account Type : 805306369                                                 |
| CN : UFC-DBPROD                                                              |
| Instance Type : 4                                                            |
| Object Class : top, person, organizationalPerson, user, computer             |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| DNS Host Name : UFC-DBProd.rt.securecompany.com                                  |
| Service Principal Name (SPN) :                                               |
|     MSSQLSvc/UFC-DBProd.rt.securecompany.com:1433                                |
|     MSSQLSvc/UFC-DBProd.rt.securecompany.com                                     |
|     TERMSRV/UFC-DBPROD                                                       |
|     TERMSRV/UFC-DBProd.rt.securecompany.com                                      |
|     WSMAN/UFC-DBProd                                                         |
|     WSMAN/UFC-DBProd.rt.securecompany.com                                        |
|     RestrictedKrbHost/UFC-DBPROD                                             |
|     HOST/UFC-DBPROD                                                          |
|     RestrictedKrbHost/UFC-DBProd.rt.securecompany.com                            |
|     HOST/UFC-DBProd.rt.securecompany.com                                         |
| Operating System : Windows Server 2016 Standard                              |
| Operating System Version : 10.0 (14393)                                      |
| UAC Flag : WORKSTATION_ACCOUNT,DONT_EXPIRE_PASSWD                            |
| User Account Control : 69632                                                 |
| Object GUID : f47d4d3c-bfa6-453e-9c2c-19d595d28e68                           |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1105                  |
| Bad Password Count : 0                                                       |
| Bad Password Time : 2019-06-17 11:37:01.6891209 +0000 UTC                    |
| When Created : 20190201064824.0Z                                             |
| When Changed : 20220419133218.0Z                                             |
| Password Last Set : 2019-06-17 11:37:13.8158567 +0000 UTC                    |
| Last Log On : 2022-04-27 15:17:05.6981141 +0000 UTC                          |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                 |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                  |
| uSNChanged : 1601-01-01 00:00:00.0646412 +0000 UTC                           |
| uSNCreated : 1601-01-01 00:00:00.0012982 +0000 UTC                           |
+------------------------------------------------------------------------------+
| DN : CN=UFC-SQLDEV,OU=Servers,DC=us,DC=funcorp,DC=local                      |
| SAM Account Name : UFC-SQLDEV$                                               |
| SAM Account Type : 805306369                                                 |
| CN : UFC-SQLDEV                                                              |
| Instance Type : 4                                                            |
| Object Class : top, person, organizationalPerson, user, computer             |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| DNS Host Name : UFC-SQLDev.rt.securecompany.com                                  |
| Service Principal Name (SPN) :                                               |
|     MSSQLSvc/UFC-SQLDev.rt.securecompany.com:1433                                |
|     MSSQLSvc/UFC-SQLDev.rt.securecompany.com                                     |
|     WSMAN/UFC-SQLDev                                                         |
|     WSMAN/UFC-SQLDev.rt.securecompany.com                                        |
|     TERMSRV/UFC-SQLDEV                                                       |
|     TERMSRV/UFC-SQLDev.rt.securecompany.com                                      |
|     RestrictedKrbHost/UFC-SQLDEV                                             |
|     HOST/UFC-SQLDEV                                                          |
|     RestrictedKrbHost/UFC-SQLDev.rt.securecompany.com                            |
|     HOST/UFC-SQLDev.rt.securecompany.com                                         |
| Operating System : Windows Server 2016 Standard                              |
| Operating System Version : 10.0 (14393)                                      |
| UAC Flag : WORKSTATION_ACCOUNT,DONT_EXPIRE_PASSWD                            |
| User Account Control : 69632                                                 |
| Object GUID : 19f7e5f0-3753-43ae-87a3-a19fc7e92247                           |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1106                  |
| Bad Password Count : 0                                                       |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                            |
| When Created : 20190201071830.0Z                                             |
| When Changed : 20220419133235.0Z                                             |
| Password Last Set : 2019-06-17 11:00:32.1796335 +0000 UTC                    |
| Last Log On : 2022-04-27 15:17:03.2634443 +0000 UTC                          |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                 |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                  |
| uSNChanged : 1601-01-01 00:00:00.0646442 +0000 UTC                           |
| uSNCreated : 1601-01-01 00:00:00.0013204 +0000 UTC                           |
+------------------------------------------------------------------------------+
```

To list all attributes of specific MS-SQL server,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command ms-sql -name ufc-sqldev 
+------------------------------------------------------------------------------+
| MS-SQL SERVERS                                                               |
+------------------------------------------------------------------------------+
| DN : CN=UFC-SQLDEV,OU=Servers,DC=us,DC=funcorp,DC=local                      |
| SAM Account Name : UFC-SQLDEV$                                               |
| SAM Account Type : 805306369                                                 |
| CN : UFC-SQLDEV                                                              |
| Instance Type : 4                                                            |
| Object Class : top, person, organizationalPerson, user, computer             |
| Object Category : CN=Computer,CN=Schema,CN=Configuration,DC=funcorp,DC=local |
| DNS Host Name : UFC-SQLDev.rt.securecompany.com                                  |
| Service Principal Name (SPN) :                                               |
|     MSSQLSvc/UFC-SQLDev.rt.securecompany.com:1433                                |
|     MSSQLSvc/UFC-SQLDev.rt.securecompany.com                                     |
|     WSMAN/UFC-SQLDev                                                         |
|     WSMAN/UFC-SQLDev.rt.securecompany.com                                        |
|     TERMSRV/UFC-SQLDEV                                                       |
|     TERMSRV/UFC-SQLDev.rt.securecompany.com                                      |
|     RestrictedKrbHost/UFC-SQLDEV                                             |
|     HOST/UFC-SQLDEV                                                          |
|     RestrictedKrbHost/UFC-SQLDev.rt.securecompany.com                            |
|     HOST/UFC-SQLDev.rt.securecompany.com                                         |
| Operating System : Windows Server 2016 Standard                              |
| Operating System Version : 10.0 (14393)                                      |
| UAC Flag : WORKSTATION_ACCOUNT,DONT_EXPIRE_PASSWD                            |
| User Account Control : 69632                                                 |
| Object GUID : 19f7e5f0-3753-43ae-87a3-a19fc7e92247                           |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1106                  |
| Bad Password Count : 0                                                       |
| Bad Password Time : 1601-01-01 00:00:00 +0000 UTC                            |
| When Created : 20190201071830.0Z                                             |
| When Changed : 20220419133235.0Z                                             |
| Password Last Set : 2019-06-17 11:00:32.1796335 +0000 UTC                    |
| Last Log On : 2022-04-27 15:19:11.3390535 +0000 UTC                          |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                 |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                  |
| uSNChanged : 1601-01-01 00:00:00.0646442 +0000 UTC                           |
| uSNCreated : 1601-01-01 00:00:00.0013204 +0000 UTC                           |
+------------------------------------------------------------------------------+
```

To list all unconstrained delegated objects,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command unconstrained 
+--------------------------+
| UNCONSTRAINED DELEGATION |
+--------------------------+
|  - UFC-DC1$              |
|  - UFC-WEBPROD$          |
+--------------------------+
```

To list SPNs available in the domain,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command spn 
+-------------------------+
| SERVICE PRINCIPAL NAMES |
+-------------------------+
|  - db1user              |
|  - dbservice            |
+-------------------------+
```

To list all attributes of specific SPN,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command spn -name dbservice
+-----------------------------------------------------------------------------------------+
| SERVICE PRINCIPAL NAMES                                                                 |
+-----------------------------------------------------------------------------------------+
| DN : CN=dbservice,CN=Users,DC=us,DC=funcorp,DC=local                                    |
| SAM Account Name : dbservice                                                            |
| SAM Account Type : 805306368                                                            |
| CN : dbservice                                                                          |
| Given Name : db                                                                         |
| SN : service                                                                            |
| Description : Account to be used for running database services which need precise time. |
| Instance Type : 4                                                                       |
| Object Class : top, person, organizationalPerson, user                                  |
| Object Category : CN=Person,CN=Schema,CN=Configuration,DC=funcorp,DC=local              |
| Service Principal Name (SPN) :                                                          |
|     TIME/UFC-DB1.rt.securecompany.com                                                       |
|     TIME/ufc-db1                                                                        |
| UAC Flag : NORMAL_ACCOUNT,DONT_EXPIRE_PASSWD,TRUSTED_TO_AUTH_FOR_DELEGATION             |
| User Account Control : 16843264                                                         |
| Object GUID : 8812bef0-c61b-47d9-9085-b11e5ea71bdb                                      |
| Object SID : S-1-5-21-3965405831-1015596948-2589850225-1120                             |
| Bad Password Count : 0                                                                  |
| Bad Password Time : 2019-02-04 10:31:05.7002426 +0000 UTC                               |
| When Created : 20190204063649.0Z                                                        |
| When Changed : 20191215055906.0Z                                                        |
| Password Last Set : 2019-02-04 06:36:49.2525007 +0000 UTC                               |
| Last Log On : 2019-12-15 05:59:06.4170143 +0000 UTC                                     |
| Last Log Off : 1601-01-01 00:00:00 +0000 UTC                                            |
| Account Expires On : 30828-09-14 02:48:05.4775807 +0000 UTC                             |
| uSNChanged : 1601-01-01 00:00:00.0504189 +0000 UTC                                      |
| uSNCreated : 1601-01-01 00:00:00.0066553 +0000 UTC                                      |
+-----------------------------------------------------------------------------------------+
```

To list AD objects in domain with highest privileges,

```c
PS C:\Users\redteamer\Desktop\shared> .\ADReaper.exe -dc rt-dc.rt.securecompany.com -user redteamer -password <password> -command admin-priv 
+---------------------------------+
| ADMIN PRIV                      |
+---------------------------------+
|  - Read-only Domain Controllers |
|  - servicesadmin                |
|  - Administrator                |
|  - krbtgt                       |
|  - Domain Controllers           |
|  - Domain Admins                |
|  - Server Operators             |
|  - Account Operators            |
|  - Administrators               |
|  - Print Operators              |
|  - Backup Operators             |
|  - Replicator                   |
+---------------------------------+
```