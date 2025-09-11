---
title: React to Hooks (MobX)
tags: [react, migration, complex, hidden]
---

This is an alternative version of the React to Hooks pattern that uses MobX.


```grit
engine marzano(0.1)
language js

// Most of the logic for this pattern is in react_hooks.grit
// https://github.com/getgrit/js/blob/main/.grit/patterns/react_hooks.grit

pattern special_first_step() {
	$use_ref_from = `useRefFrom`,
	// Avoid inserting the "Handler" suffix
	$handler_callback_suffix = .,
	first_step($use_ref_from, $handler_callback_suffix)
}

sequential {
	file(body=program(statements=some bubble($program) special_first_step())),
	// Run it 3 times to converge
	file(body=second_step(handler_callback_suffix=.)),
	file(body=second_step(handler_callback_suffix=.)),
	file(body=second_step(handler_callback_suffix=.)),
	file($body) where {
		$body <: program($statements),
		$statements <: bubble($body, $program) and { maybe adjust_imports(),
		add_more_imports(use_ref_from=`"~/hooks/myhooks"`) }
	}
}
```

## Respects `useRefFrom`

```js
import React from "react";
import styled from "styled-components";

import { CustomComponent } from "components/CustomComponent/CustomComponent";
import { IBrand } from "models/brand";
import { Banner, IBannerProps } from "models/viewport";
import { BannerPicture } from "models/banner";

export interface IMainProps {
  bannerStuff: IBannerProps;
  dataHeaderRef?: React.RefCallback<HTMLElement>;
}

class BrandHeaderBase extends React.Component<
  IMainProps & IBrandProps
> {
  render() {
    const {
      bannerStuff,
      dataHeaderRef,
      brandName
    } = this.props;
    return (
      <Banner>
        <h3>Some text</h3>
        <p>Some more text</p>
        {this.renderBannerDetails()}
      </Banner>
    );
  }

  private name = "BrandHeader";

  private invoker: React.RefObject<HTMLElement> = React.createRef();
  private util: number = 9;

  private renderBannerDetails = () => {
    if (!getGoodStuff()) {
      return this.props.viewport.isMedium ? (
        <InternalBrand
            brand={brand}
            height={240}
          />
      ) : null;
    } else {
      const CustomBanner: React.FC<{ height: number }> = ({
        height,
      }) => (
        <InternalBrand
            brand={brand}
            height={height}
          />
      );

      return (
        <ResponsiveBanner>
          <CustomBanner height={240} />
        </ResponsiveBanner>
      );
    }
  };
}
```

```js
import { useRefFrom } from '~/hooks/myhooks';
import React, { useRef } from 'react';
import styled from 'styled-components';

import { CustomComponent } from 'components/CustomComponent/CustomComponent';
import { IBrand } from 'models/brand';
import { Banner, IBannerProps } from 'models/viewport';
import { BannerPicture } from 'models/banner';

export interface IMainProps {
  bannerStuff: IBannerProps;
  dataHeaderRef?: React.RefCallback<HTMLElement>;
}

const BrandHeaderBase: React.FunctionComponent<IMainProps & IBrandProps> = (props) => {
  const name = useRefFrom(() => 'BrandHeader').current;
  const invoker = useRef<React.RefObject<HTMLElement>>();
  const util = useRefFrom(() => 9).current;
  const renderBannerDetails = () => {
    if (!getGoodStuff()) {
      return props.viewport.isMedium ? <InternalBrand brand={brand} height={240} /> : null;
    } else {
      const CustomBanner: React.FC<{ height: number }> = ({ height }) => (
        <InternalBrand brand={brand} height={height} />
      );

      return (
        <ResponsiveBanner>
          <CustomBanner height={240} />
        </ResponsiveBanner>
      );
    }
  };

  const { bannerStuff, dataHeaderRef, brandName } = props;
  return (
    <Banner>
      <h3>Some text</h3>
      <p>Some more text</p>
      {renderBannerDetails()}
    </Banner>
  );
};
```
