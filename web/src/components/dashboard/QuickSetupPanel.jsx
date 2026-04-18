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

import React, { useContext, useState } from 'react';
import { Card, Typography, Button } from '@douyinfe/semi-ui';
import { Terminal, Copy, Check, ChevronDown, ChevronRight } from 'lucide-react';
import { StatusContext } from '../../context/Status';
import { UserContext } from '../../context/User';

const { Text, Title } = Typography;

const QuickSetupPanel = ({ CARD_PROPS, FLEX_CENTER_GAP2, t }) => {
  const [statusState] = useContext(StatusContext);
  const [userState] = useContext(UserContext);
  const [copied, setCopied] = useState(null);

  const serverAddress = statusState?.status?.server_address || '';
  const apiKey = userState?.user?.access_token || 'sk-your-api-key';

  const npxCommand = `npx openroutera-setup --key=${apiKey}`;

  const manualConfig = JSON.stringify(
    {
      env: {
        ANTHROPIC_BASE_URL: serverAddress,
        ANTHROPIC_API_KEY: apiKey,
        API_TIMEOUT_MS: '3000000',
        CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC: 1,
        ANTHROPIC_DEFAULT_SONNET_MODEL: 'glm-4.7',
        ANTHROPIC_DEFAULT_OPUS_MODEL: 'glm-5.1',
      },
    },
    null,
    2,
  );

  const handleCopy = (text, key) => {
    navigator.clipboard.writeText(text).then(() => {
      setCopied(key);
      setTimeout(() => setCopied(null), 2000);
    });
  };

  const CopyButton = ({ text, copyKey }) => (
    <button
      onClick={() => handleCopy(text, copyKey)}
      className='inline-flex items-center gap-1 px-2 py-1 text-xs rounded-md bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600 transition-colors'
      style={{ border: 'none', cursor: 'pointer' }}
    >
      {copied === copyKey ? (
        <>
          <Check size={12} className='text-green-500' />
          <span className='text-green-600 dark:text-green-400'>
            {t('已复制')}
          </span>
        </>
      ) : (
        <>
          <Copy size={12} />
          <span>{t('复制')}</span>
        </>
      )}
    </button>
  );

  return (
    <Card
      {...CARD_PROPS}
      className='bg-gray-50 border-0 !rounded-2xl'
      title={
        <div className={FLEX_CENTER_GAP2}>
          <Terminal size={16} />
          {t('快速配置')}
        </div>
      }
    >
      <div className='space-y-4'>
        {/* Method 1: NPX */}
        <div>
          <div className='flex items-center justify-between mb-2'>
            <Text strong>{t('一键配置 Claude Code')}</Text>
          </div>
          <div className='bg-gray-900 rounded-lg p-3 relative group'>
            <code className='text-sm text-green-400 break-all'>
              {npxCommand}
            </code>
            <div className='absolute top-2 right-2'>
              <CopyButton text={npxCommand} copyKey='npx' />
            </div>
          </div>
          <Text
            type='tertiary'
            size='small'
            className='block mt-1'
          >
            {t('在终端运行此命令，自动完成 Claude Code 配置')}
          </Text>
        </div>

        {/* Method 2: Manual */}
        <div>
          <details className='group'>
            <summary className='flex items-center gap-2 cursor-pointer text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100'>
              <ChevronRight
                size={14}
                className='transition-transform group-open:rotate-90'
              />
              {t('手动配置')}
            </summary>
            <div className='mt-2 space-y-2'>
              <Text type='tertiary' size='small'>
                {t('编辑 ~/.claude/settings.json，添加以下内容：')}
              </Text>
              <div className='bg-gray-900 rounded-lg p-3 relative'>
                <pre className='text-sm text-gray-300 overflow-x-auto'>
                  <code>{manualConfig}</code>
                </pre>
                <div className='absolute top-2 right-2'>
                  <CopyButton text={manualConfig} copyKey='manual' />
                </div>
              </div>
            </div>
          </details>
        </div>

        {/* Env vars reference */}
        <div>
          <details className='group'>
            <summary className='flex items-center gap-2 cursor-pointer text-sm font-medium text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100'>
              <ChevronRight
                size={14}
                className='transition-transform group-open:rotate-90'
              />
              {t('环境变量说明')}
            </summary>
            <div className='mt-2 space-y-1'>
              {[
                ['ANTHROPIC_BASE_URL', t('API 服务地址')],
                ['ANTHROPIC_API_KEY', t('你的 API 密钥')],
                ['ANTHROPIC_DEFAULT_SONNET_MODEL', t('默认 Sonnet 模型')],
                ['ANTHROPIC_DEFAULT_OPUS_MODEL', t('默认 Opus 模型')],
                ['API_TIMEOUT_MS', t('请求超时时间 (ms)')],
              ].map(([key, desc]) => (
                <div key={key} className='flex items-start gap-2 text-xs'>
                  <code className='text-blue-600 dark:text-blue-400 whitespace-nowrap font-mono'>
                    {key}
                  </code>
                  <span className='text-gray-500'>{desc}</span>
                </div>
              ))}
            </div>
          </details>
        </div>
      </div>
    </Card>
  );
};

export default QuickSetupPanel;
