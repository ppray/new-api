/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/

import React, { useEffect, useState, useRef } from 'react';
import { Banner, Button, Form, Row, Col, Spin } from '@douyinfe/semi-ui';
import {
  API,
  removeTrailingSlash,
  showError,
  showSuccess,
} from '../../../helpers';
import { useTranslation } from 'react-i18next';
import { BookOpen, TriangleAlert } from 'lucide-react';

export default function SettingsPaymentGatewayNowPayments(props) {
  const { t } = useTranslation();
  const sectionTitle = props.hideSectionTitle
    ? undefined
    : t('NowPayments 设置');
  const [loading, setLoading] = useState(false);
  const [inputs, setInputs] = useState({
    NowPaymentsEnabled: false,
    NowPaymentsApiKey: '',
    NowPaymentsWebhookSecret: '',
    NowPaymentsCurrency: 'usdtsol',
    NowPaymentsMinTopUp: 1,
    NowPaymentsUnitPrice: 1.0,
  });
  const [originInputs, setOriginInputs] = useState({});
  const formApiRef = useRef(null);

  useEffect(() => {
    if (props.options && formApiRef.current) {
      const currentInputs = {
        NowPaymentsEnabled:
          props.options.NowPaymentsEnabled !== undefined
            ? props.options.NowPaymentsEnabled
            : false,
        NowPaymentsApiKey: props.options.NowPaymentsApiKey || '',
        NowPaymentsWebhookSecret:
          props.options.NowPaymentsWebhookSecret || '',
        NowPaymentsCurrency:
          props.options.NowPaymentsCurrency || 'usdtsol',
        NowPaymentsMinTopUp:
          props.options.NowPaymentsMinTopUp !== undefined
            ? parseFloat(props.options.NowPaymentsMinTopUp)
            : 1,
        NowPaymentsUnitPrice:
          props.options.NowPaymentsUnitPrice !== undefined
            ? parseFloat(props.options.NowPaymentsUnitPrice)
            : 1.0,
      };
      setInputs(currentInputs);
      setOriginInputs({ ...currentInputs });
      formApiRef.current.setValues(currentInputs);
    }
  }, [props.options]);

  const handleFormChange = (values) => {
    setInputs(values);
  };

  const submitNowPaymentsSetting = async () => {
    if (props.options.ServerAddress === '') {
      showError(t('请先填写服务器地址'));
      return;
    }

    setLoading(true);
    try {
      const options = [];

      if (
        originInputs['NowPaymentsEnabled'] !== inputs.NowPaymentsEnabled &&
        inputs.NowPaymentsEnabled !== undefined
      ) {
        options.push({
          key: 'NowPaymentsEnabled',
          value: inputs.NowPaymentsEnabled ? 'true' : 'false',
        });
      }

      if (
        inputs.NowPaymentsApiKey &&
        inputs.NowPaymentsApiKey !== ''
      ) {
        options.push({
          key: 'NowPaymentsApiKey',
          value: inputs.NowPaymentsApiKey,
        });
      }

      if (
        inputs.NowPaymentsWebhookSecret &&
        inputs.NowPaymentsWebhookSecret !== ''
      ) {
        options.push({
          key: 'NowPaymentsWebhookSecret',
          value: inputs.NowPaymentsWebhookSecret,
        });
      }

      if (inputs.NowPaymentsCurrency !== '') {
        options.push({
          key: 'NowPaymentsCurrency',
          value: inputs.NowPaymentsCurrency,
        });
      }

      if (
        inputs.NowPaymentsUnitPrice !== undefined &&
        inputs.NowPaymentsUnitPrice !== null
      ) {
        options.push({
          key: 'NowPaymentsUnitPrice',
          value: inputs.NowPaymentsUnitPrice.toString(),
        });
      }

      if (
        inputs.NowPaymentsMinTopUp !== undefined &&
        inputs.NowPaymentsMinTopUp !== null
      ) {
        options.push({
          key: 'NowPaymentsMinTopUp',
          value: inputs.NowPaymentsMinTopUp.toString(),
        });
      }

      const requestQueue = options.map((opt) =>
        API.put('/api/option/', {
          key: opt.key,
          value: opt.value,
        }),
      );

      const results = await Promise.all(requestQueue);

      const errorResults = results.filter((res) => !res.data.success);
      if (errorResults.length > 0) {
        errorResults.forEach((res) => {
          showError(res.data.message);
        });
      } else {
        showSuccess(t('更新成功'));
        setOriginInputs({ ...inputs });
        props.refresh?.();
      }
    } catch (error) {
      showError(t('更新失败'));
    }
    setLoading(false);
  };

  return (
    <Spin spinning={loading}>
      <Form
        initValues={inputs}
        onValueChange={handleFormChange}
        getFormApi={(api) => (formApiRef.current = api)}
      >
        <Form.Section text={sectionTitle}>
          <Banner
            type='info'
            icon={<BookOpen size={16} />}
            description={
              <>
                NowPayments 密钥、Webhook 等设置请
                <a
                  href='https://account.nowpayments.io/'
                  target='_blank'
                  rel='noreferrer'
                >
                  点击此处
                </a>
                进行设置。
                <br />
                {t('回调地址')}：
                {props.options.ServerAddress
                  ? removeTrailingSlash(props.options.ServerAddress)
                  : t('网站地址')}
                /api/nowpayments/webhook
              </>
            }
            style={{ marginBottom: 12 }}
          />
          <Banner
            type='warning'
            icon={<TriangleAlert size={16} />}
            description={t(
              '请确保在 NowPayments 后台配置了相同的 IPN 密钥，否则回调验签会失败。',
            )}
            style={{ marginBottom: 16 }}
          />

          <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 24, xl: 24, xxl: 24 }}>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Switch
                field='NowPaymentsEnabled'
                label={t('启用 NowPayments')}
                size='default'
                checkedText='｜'
                uncheckedText='〇'
              />
            </Col>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Input
                field='NowPaymentsCurrency'
                label={t('支付货币')}
                placeholder='usdtsol'
                extraText={t('NowPayments 货币代码，例如 usdtsol（Solana USDT）、usdcsol（Solana USDC）')}
              />
            </Col>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.InputNumber
                field='NowPaymentsMinTopUp'
                label={t('最低充值美元数量')}
                placeholder={t('例如：2，就是最低充值2$')}
                extraText={t('用户单次最少可充值的美元数量')}
                min={1}
              />
            </Col>
          </Row>

          <Row
            gutter={{ xs: 8, sm: 16, md: 24, lg: 24, xl: 24, xxl: 24 }}
            style={{ marginTop: 16 }}
          >
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Input
                field='NowPaymentsApiKey'
                label={t('API 密钥')}
                placeholder={t('例如：XXX-XXX，留空表示保持当前不变')}
                extraText={t(
                  '保存后不会回显，请填写当前环境对应的 NowPayments API 密钥',
                )}
                type='password'
              />
            </Col>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.Input
                field='NowPaymentsWebhookSecret'
                label={t('IPN 密钥')}
                placeholder={t('例如：XXX，留空表示保持当前不变')}
                extraText={t(
                  '用于校验 NowPayments Webhook 签名，保存后不会回显',
                )}
                type='password'
              />
            </Col>
            <Col xs={24} sm={24} md={8} lg={8} xl={8}>
              <Form.InputNumber
                field='NowPaymentsUnitPrice'
                precision={2}
                label={t('充值价格（x元/美金）')}
                placeholder={t('例如：7，就是7元/美金')}
                extraText={t('按 1 美元对应的站内价格填写')}
                min={0}
              />
            </Col>
          </Row>

          <Button
            onClick={submitNowPaymentsSetting}
            style={{ marginTop: 16 }}
          >
            {t('更新 NowPayments 设置')}
          </Button>
        </Form.Section>
      </Form>
    </Spin>
  );
}
