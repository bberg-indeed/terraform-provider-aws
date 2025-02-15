---
subcategory: "ELB (Elastic Load Balancing)"
layout: "aws"
page_title: "AWS: aws_lb_listener"
description: |-
  Provides a Load Balancer Listener resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lb_listener

Provides a Load Balancer Listener resource.

~> **Note:** `aws_alb_listener` is known as `aws_lb_listener`. The functionality is identical.

## Example Usage

### Forward Action

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
import { LbTargetGroup } from "./.gen/providers/aws/lb-target-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const frontEnd = new Lb(this, "front_end", {});
    const awsLbTargetGroupFrontEnd = new LbTargetGroup(this, "front_end_1", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbTargetGroupFrontEnd.overrideLogicalId("front_end");
    const awsLbListenerFrontEnd = new LbListener(this, "front_end_2", {
      certificateArn:
        "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
      defaultAction: [
        {
          targetGroupArn: Token.asString(awsLbTargetGroupFrontEnd.arn),
          type: "forward",
        },
      ],
      loadBalancerArn: frontEnd.arn,
      port: Token.asNumber("443"),
      protocol: "HTTPS",
      sslPolicy: "ELBSecurityPolicy-2016-08",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerFrontEnd.overrideLogicalId("front_end");
  }
}

```

To a NLB:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { LbListener } from "./.gen/providers/aws/lb-listener";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new LbListener(this, "front_end", {
      alpnPolicy: "HTTP2Preferred",
      certificateArn:
        "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
      defaultAction: [
        {
          targetGroupArn: Token.asString(awsLbTargetGroupFrontEnd.arn),
          type: "forward",
        },
      ],
      loadBalancerArn: Token.asString(awsLbFrontEnd.arn),
      port: Token.asNumber("443"),
      protocol: "TLS",
    });
  }
}

```

### Redirect Action

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const frontEnd = new Lb(this, "front_end", {});
    const awsLbListenerFrontEnd = new LbListener(this, "front_end_1", {
      defaultAction: [
        {
          redirect: {
            port: "443",
            protocol: "HTTPS",
            statusCode: "HTTP_301",
          },
          type: "redirect",
        },
      ],
      loadBalancerArn: frontEnd.arn,
      port: Token.asNumber("80"),
      protocol: "HTTP",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerFrontEnd.overrideLogicalId("front_end");
  }
}

```

### Fixed-response Action

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const frontEnd = new Lb(this, "front_end", {});
    const awsLbListenerFrontEnd = new LbListener(this, "front_end_1", {
      defaultAction: [
        {
          fixedResponse: {
            contentType: "text/plain",
            messageBody: "Fixed response content",
            statusCode: "200",
          },
          type: "fixed-response",
        },
      ],
      loadBalancerArn: frontEnd.arn,
      port: Token.asNumber("80"),
      protocol: "HTTP",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerFrontEnd.overrideLogicalId("front_end");
  }
}

```

### Authenticate-cognito Action

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CognitoUserPool } from "./.gen/providers/aws/cognito-user-pool";
import { CognitoUserPoolClient } from "./.gen/providers/aws/cognito-user-pool-client";
import { CognitoUserPoolDomain } from "./.gen/providers/aws/cognito-user-pool-domain";
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
import { LbTargetGroup } from "./.gen/providers/aws/lb-target-group";
interface MyConfig {
  name: any;
  name1: any;
  userPoolId: any;
  domain: any;
  userPoolId1: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    const pool = new CognitoUserPool(this, "pool", {
      name: config.name,
    });
    const client = new CognitoUserPoolClient(this, "client", {
      name: config.name1,
      userPoolId: config.userPoolId,
    });
    const domain = new CognitoUserPoolDomain(this, "domain", {
      domain: config.domain,
      userPoolId: config.userPoolId1,
    });
    const frontEnd = new Lb(this, "front_end", {});
    const awsLbTargetGroupFrontEnd = new LbTargetGroup(this, "front_end_4", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbTargetGroupFrontEnd.overrideLogicalId("front_end");
    const awsLbListenerFrontEnd = new LbListener(this, "front_end_5", {
      defaultAction: [
        {
          authenticateCognito: {
            userPoolArn: pool.arn,
            userPoolClientId: client.id,
            userPoolDomain: domain.domain,
          },
          type: "authenticate-cognito",
        },
        {
          targetGroupArn: Token.asString(awsLbTargetGroupFrontEnd.arn),
          type: "forward",
        },
      ],
      loadBalancerArn: frontEnd.arn,
      port: Token.asNumber("80"),
      protocol: "HTTP",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerFrontEnd.overrideLogicalId("front_end");
  }
}

```

### Authenticate-OIDC Action

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
import { LbTargetGroup } from "./.gen/providers/aws/lb-target-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const frontEnd = new Lb(this, "front_end", {});
    const awsLbTargetGroupFrontEnd = new LbTargetGroup(this, "front_end_1", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbTargetGroupFrontEnd.overrideLogicalId("front_end");
    const awsLbListenerFrontEnd = new LbListener(this, "front_end_2", {
      defaultAction: [
        {
          authenticateOidc: {
            authorizationEndpoint: "https://example.com/authorization_endpoint",
            clientId: "client_id",
            clientSecret: "client_secret",
            issuer: "https://example.com",
            tokenEndpoint: "https://example.com/token_endpoint",
            userInfoEndpoint: "https://example.com/user_info_endpoint",
          },
          type: "authenticate-oidc",
        },
        {
          targetGroupArn: Token.asString(awsLbTargetGroupFrontEnd.arn),
          type: "forward",
        },
      ],
      loadBalancerArn: frontEnd.arn,
      port: Token.asNumber("80"),
      protocol: "HTTP",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerFrontEnd.overrideLogicalId("front_end");
  }
}

```

### Gateway Load Balancer Listener

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
import { LbTargetGroup } from "./.gen/providers/aws/lb-target-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new Lb(this, "example", {
      loadBalancerType: "gateway",
      name: "example",
      subnetMapping: [
        {
          subnetId: Token.asString(awsSubnetExample.id),
        },
      ],
    });
    const awsLbTargetGroupExample = new LbTargetGroup(this, "example_1", {
      healthCheck: {
        port: Token.asString(80),
        protocol: "HTTP",
      },
      name: "example",
      port: 6081,
      protocol: "GENEVE",
      vpcId: Token.asString(awsVpcExample.id),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbTargetGroupExample.overrideLogicalId("example");
    const awsLbListenerExample = new LbListener(this, "example_2", {
      defaultAction: [
        {
          targetGroupArn: Token.asString(awsLbTargetGroupExample.id),
          type: "forward",
        },
      ],
      loadBalancerArn: example.id,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerExample.overrideLogicalId("example");
  }
}

```

### Mutual TLS Authentication

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Lb } from "./.gen/providers/aws/lb";
import { LbListener } from "./.gen/providers/aws/lb-listener";
import { LbTargetGroup } from "./.gen/providers/aws/lb-target-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new Lb(this, "example", {
      loadBalancerType: "application",
    });
    const awsLbTargetGroupExample = new LbTargetGroup(this, "example_1", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbTargetGroupExample.overrideLogicalId("example");
    const awsLbListenerExample = new LbListener(this, "example_2", {
      defaultAction: [
        {
          targetGroupArn: Token.asString(awsLbTargetGroupExample.id),
          type: "forward",
        },
      ],
      loadBalancerArn: example.id,
      mutualAuthentication: {
        mode: "verify",
        trustStoreArn: "...",
      },
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsLbListenerExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `defaultAction` - (Required) Configuration block for default actions. Detailed below.
* `loadBalancerArn` - (Required, Forces New Resource) ARN of the load balancer.

The following arguments are optional:

* `alpnPolicy` - (Optional)  Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
* `certificateArn` - (Optional) ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](/docs/providers/aws/r/lb_listener_certificate.html).
* `mutualAuthentication` - (Optional) The mutual authentication configuration information. Detailed below.
* `port` - (Optional) Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
* `protocol` - (Optional) Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
* `sslPolicy` - (Optional) Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

~> **NOTE::** Please note that listeners that are attached to Application Load Balancers must use either `HTTP` or `HTTPS` protocols while listeners that are attached to Network Load Balancers must use the `TCP` protocol.

### default_action

The following arguments are required:

* `type` - (Required) Type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.

The following arguments are optional:

* `authenticateCognito` - (Optional) Configuration block for using Amazon Cognito to authenticate users. Specify only when `type` is `authenticate-cognito`. Detailed below.
* `authenticateOidc` - (Optional) Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when `type` is `authenticate-oidc`. Detailed below.
* `fixedResponse` - (Optional) Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
* `forward` - (Optional) Configuration block for creating an action that distributes requests among one or more target groups. Specify only if `type` is `forward`. If you specify both `forward` block and `targetGroupArn` attribute, you can specify only one target group using `forward` and it must be the same target group specified in `targetGroupArn`. Detailed below.
* `order` - (Optional) Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between `1` and `50000`.
* `redirect` - (Optional) Configuration block for creating a redirect action. Required if `type` is `redirect`. Detailed below.
* `targetGroupArn` - (Optional) ARN of the Target Group to which to route traffic. Specify only if `type` is `forward` and you want to route to a single target group. To route to one or more target groups, use a `forward` block instead.

#### authenticate_cognito

The following arguments are required:

* `userPoolArn` - (Required) ARN of the Cognito user pool.
* `userPoolClientId` - (Required) ID of the Cognito user pool client.
* `userPoolDomain` - (Required) Domain prefix or fully-qualified domain name of the Cognito user pool.

The following arguments are optional:

* `authenticationRequestExtraParams` - (Optional) Query parameters to include in the redirect request to the authorization endpoint. Max: 10. Detailed below.
* `onUnauthenticatedRequest` - (Optional) Behavior if the user is not authenticated. Valid values are `deny`, `allow` and `authenticate`.
* `scope` - (Optional) Set of user claims to be requested from the IdP.
* `sessionCookieName` - (Optional) Name of the cookie used to maintain session information.
* `sessionTimeout` - (Optional) Maximum duration of the authentication session, in seconds.

##### authentication_request_extra_params

* `key` - (Required) Key of query parameter.
* `value` - (Required) Value of query parameter.

#### authenticate_oidc

The following arguments are required:

* `authorizationEndpoint` - (Required) Authorization endpoint of the IdP.
* `clientId` - (Required) OAuth 2.0 client identifier.
* `clientSecret` - (Required) OAuth 2.0 client secret.
* `issuer` - (Required) OIDC issuer identifier of the IdP.
* `tokenEndpoint` - (Required) Token endpoint of the IdP.
* `userInfoEndpoint` - (Required) User info endpoint of the IdP.

The following arguments are optional:

* `authenticationRequestExtraParams` - (Optional) Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
* `onUnauthenticatedRequest` - (Optional) Behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
* `scope` - (Optional) Set of user claims to be requested from the IdP.
* `sessionCookieName` - (Optional) Name of the cookie used to maintain session information.
* `sessionTimeout` - (Optional) Maximum duration of the authentication session, in seconds.

#### fixed_response

The following arguments are required:

* `contentType` - (Required) Content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.

The following arguments are optional:

* `messageBody` - (Optional) Message body.
* `statusCode` - (Optional) HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

#### forward

The following arguments are required:

* `targetGroup` - (Required) Set of 1-5 target group blocks. Detailed below.

The following arguments are optional:

* `stickiness` - (Optional) Configuration block for target group stickiness for the rule. Detailed below.

##### target_group

The following arguments are required:

* `arn` - (Required) ARN of the target group.

The following arguments are optional:

* `weight` - (Optional) Weight. The range is 0 to 999.

##### stickiness

The following arguments are required:

* `duration` - (Required) Time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).

The following arguments are optional:

* `enabled` - (Optional) Whether target group stickiness is enabled. Default is `false`.

#### redirect

~> **NOTE::** You can reuse URI components using the following reserved keywords: `#{protocol}`, `#{host}`, `#{port}`, `#{path}` (the leading "/" is removed) and `#{query}`.

The following arguments are required:

* `statusCode` - (Required) HTTP redirect code. The redirect is either permanent (`HTTP_301`) or temporary (`HTTP_302`).

The following arguments are optional:

* `host` - (Optional) Hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
* `path` - (Optional) Absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
* `port` - (Optional) Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
* `protocol` - (Optional) Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
* `query` - (Optional) Query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.

### mutual_authentication

* `mode` - (Required) Valid values are `off`, `verify` and `passthrough`.
* `trustStoreArn` - (Required) ARN of the elbv2 Trust Store.
* `ignoreClientCertificateExpiry` - (Optional) Whether client certificate expiry is ignored. Default is `false`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the listener (matches `id`).
* `id` - ARN of the listener (matches `arn`).
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import listeners using their ARN. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
  }
}

```

Using `terraform import`, import listeners using their ARN. For example:

```console
% terraform import aws_lb_listener.front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
```

<!-- cache-key: cdktf-0.20.0 input-174713c7c737454764e5033a2fcb15ce31c91908eaf73c511335922f57b68249 -->