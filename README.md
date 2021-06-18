This Terraform provider enables create, read, update, delete, and import operations for Expensify policy users. It also enables create, read, and import operations for Policy.


## Requirements

* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* Application: [Expensify](https://www.expensify.com/) (API is supported in collect and control policy plans)
* [Expensify API Documentation](https://integrations.expensify.com/Integration-Server/doc/employeeUpdater.html)


## Application Account

### Setup<a id="setup"></a>
1. Create an expensify account at https://www.expensify.com/<br>
2. To create a policy, go to `Settings -> Policies -> Group -> click on New Policy`.<br>
3. After creating the policy, for policy ID, go to `Settings -> Policies -> Group -> Select the appropriate policy` and note the policy ID from the URL.<br>
   For example, in Policy url - ```"https://www.expensify.com/policy?param={policyID:22E95AFCD33ABE2BB8}", "22E95AFCD33ABE2BB8" is Policy ID```

### API Authentication
 *Generate credentials from an account which is admin to both domain and policy*
1. To authenticate API, we need a pair of credentials: expensifyUserID and expensifyUserSecret.<br>
2. For this, go to https://www.expensify.com/tools/integrations/ and generate the credentials.<br>
3. A pair of credentials: expensifyUserID and expensifyUserSecret will be generated and shown on the page.<br>


## Building The Provider
1. Clone the repository, add all the dependencies and create a vendor directory that contains all dependencies. For this, run the following commands: <br>
```
cd terraform-provider-expensify
go mod init terraform-provider-expensify
go mod tidy
go mod vendor
```

## Managing terraform plugins
*For Windows:*
1. Run the following command to create a vendor sub-directory (`%APPDATA%/terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${OS_ARCH}`) which will consist of all terraform plugins. <br> 
Command: 
```bash
mkdir -p %APPDATA%/terraform.d/plugins/expensify.com/employee/expensify/1.0.0/windows_amd64
```
2. Run `go build -o terraform-provider-expensify.exe` to generate the binary in present working directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-expensify.exe %APPDATA%\terraform.d\plugins\expensify.com\employee\expensify\1.0.0\windows_amd64
 ``` 
<p align="center">[OR]</p>
 
3. Manually move the file from current directory to destination directory (`%APPDATA%\terraform.d\plugins\expensify.com\employee\expensify\1.0.0\windows_amd64`).<br>


## Working with terraform

### Application Credential Integration in terraform
1. Add `terraform` block and `provider` block as shown in [example usage](#example-usage).
2. Get a pair of credentials: expensifyUserID and expensifyUserSecret. For this, visit https://www.expensify.com/tools/integrations/.
3. Assign the above credentials to the respective field in the `provider` block.

### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the changes.

### Create User
1. Add the `employee_email`, `manager_email`, `policy_id`, `first_name`, `last_name` in the respective field in `resource` block as shown in [example usage](#example-usage).
2. Refer to [setup](#setup) for the policy ID.
3. Run the basic terraform commands.<br>
4. On successful execution, sends an account setup mail to user.<br>

### Update the user
1. Update the data of the user in the `resource` block as show in [example usage](#example-usage) and run the basic terraform commands to update user. 
   User is not allowed to update `employee_email` and `policy_id`.

### Read the User Data
Add `data` and `output` blocks as shown in the [example usage](#example-usage) and run the basic terraform commands.

### Delete the user
Delete the `resource` block of the user and run `terraform apply`.

### Import a User Data
1. Write manually a `resource` configuration block for the user as shown in [example usage](#example-usage). Imported user will be mapped to this block.
2. Run the command `terraform import expensify_user.employee [POLICY_ID]:[EMAIL_ID]` to import user.
3. Refer to [setup](#setup) for the policy ID.
4. Run `terraform plan`, if output shows `0 to addd, 0 to change and 0 to destroy` user import is successful, otherwise recheck the employee data in `resource` block with employee data in the policy in Expensify website.

### Create Policy
1. Add the `policy_name`, `plan` in the respective field in `resource` block as shown in [example usage](#example-usage).
2. Run the basic terraform commands.<br>

### Read the Policy Data
Add `data` and `output` blocks as shown in the [example usage](#example-usage) and run the basic terraform commands.

### Import a Policy Data
1. Write manually a `resource` configuration block for the policy as shown in [example usage](#example-usage). Imported policy will be mapped to this block.
2. Run the command `terraform import expensify_policy.policy [POLICY_ID]` to import policy.
3. Refer to [setup](#setup) for the policy ID.
4. Run `terraform plan`, if output shows `0 to addd, 0 to change and 0 to destroy` policy import is successful, otherwise recheck the employee data in `resource` block with policy data in Expensify website.


## Example Usage<a id="example-usage"></a>

```
terraform{
    required_providers {
        expensify = {
            version = "1.0.0"
            source = "expensify.com/employee/expensify"
        }
    }
}

provider "expensify" {
    expensify_user_id = "_REPLACE_EXPENSIFY_USER_ID_"
    expensify_user_secret = "_REPLACE_EXPENSIFY_USER_SECRET_" 
}

resource "expensify_policy" "policy"{
    policy_name = "demo"
    plan = "corporate"
}

output "resource_policy"{
    value = expensify_policy.employee
}

data "expensify_policy" "policy" {
    policy_id = "22E95AFCD33ABE2BB8"
}

output "datasouce_policy"{
    value = data.expensify_policy.employee
}

resource "expensify_user" "employee"{
    employee_email = "employee@domain.com"
    manager_email = "manager@domain.com"
    policy_id = "22E95AFCD33ABE2BB8"
    employee_id = "101"
    first_name = "Dummy"
    last_name = "Employee"
    approves_to = "approver@domain.com"
    approval_limit = 5
    over_limit_approver = "overlimitapprover@domain.com"
}

output "resource_user"{
    value = expensify_user.employee
}

data "expensify_user" "employee" {
    policy_id = "22E95AFCD33ABE2BB8"
    employee_email = "employee@domain.com" 
}

output "datasouce_user"{
    value = data.expensify_user.employee
}
```


## Argument Reference

* `expensify_user_id` (Required, String) - The Expensify expensify User ID. This may also be set via the `"EXPENSIFY_USER_ID"` environment variable.
* `expensify_user_secret` (Required, String) - The Expensify expensify User Secret. This may also be set via the `"EXPENSIFY_USER_SECRET"` environment variable.
* `employee_email` (Required, String) - The email address of the employee.
* `manager_email` (Optional, String) - Manager email address.
* `policy_id` (Required, String) - The ID of policy for which employee is to be added.
* `first_name` (Optional, String) - First name of the employee in Expensify. 
* `last_name` (Optional, String) - Last name of the employee in Expensify.
* `employee_id` (Optional, String) - Unique ID of the Employee.
* `over_limit_approver` (Optional, String) - over limit approver email address. Required if an `approval_limit` is specified.
* `approval_limit` (Optional, Float) - Specifies limit of report total.
* `approves_to` (Optional, String) - approver email address.
* `policy_name` (Required, String) - Name of the policy.
* `plan` (Optional, String) - Defines the plan for the policy. Supported values are `team` (Collect) and `corporate` (Control). Default value is `team`. 


## Exceptions

* Managing user role should be done through UI.
* Updating of the fields `manager_email`, `approves_to`, `over_limit_approver`, and `approval_limit` is meaningful only if Approval Mode for policy is Advanced Approval.
* Updating `first_name` and `last_name` in any one policy will automatically update them in other policies.
* Not allowed overwriting `first_name` and `last_name` values manually set by the employee in their Expensify account.
* To add an employee to multiple policies, write multiple `resource` block with different policy ID.
* Once the value of any attribute is set, it cannot be set back to null through provider. But, you can set it to null via UI.
