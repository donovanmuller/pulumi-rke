// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rke.Inputs
{

    public sealed class ClusterPrivateRegistryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set as default registry. Default `false` (bool)
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Registry password (string)
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Registry URL (string)
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Registry user (string)
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public ClusterPrivateRegistryArgs()
        {
        }
    }
}