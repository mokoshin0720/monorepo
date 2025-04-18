import type { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from 'App';
import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
import React, { useEffect, useState } from 'react';
import { View, StyleSheet, Text } from 'react-native';
import { Button } from 'src/components/Button';
import { API_HOST, APP_ENV } from 'src/config/env';
import { MemoModel } from 'src/model/Memo';

type MemoListScreenNavigationProp = NativeStackScreenProps<
    RootStackParamList,
    'MemoList'
>;

export const MemoListScreen: React.FC<MemoListScreenNavigationProp> = ({
    navigation,
    route,
}) => {
    const [memos, setMemos] = useState<string[]>([]);

    useEffect(() => {
        getMemos();
    }, []);

    const options: AxiosRequestConfig = {
        method: 'GET',
        url: `${API_HOST}/memos`,
    };

    const getMemos = async () => {
        axios(options)
            .then((res: AxiosResponse<MemoModel[]>) => {
                const { data, status } = res;
                console.log(status);
                console.log(data);
            })
            .catch((e: AxiosError<{ error: string }>) => {
                console.log(e.message);
            });
    };

    if (memos.length === 0) {
        return (
            <View style={emptyStyles.container}>
                <View style={emptyStyles.inner}>
                    <Text style={emptyStyles.title}>
                        最初のメモを作成しよう！
                    </Text>
                    <Text>{API_HOST}</Text>
                    <Text>{APP_ENV}</Text>
                </View>
                <Button
                    label="作成する"
                    style={emptyStyles.button}
                    onPress={() => {
                        navigation.navigate('MemoCreate');
                    }}
                />
            </View>
        );
    }
};

const emptyStyles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
    inner: {
        justifyContent: 'center',
        alignItems: 'center',
    },
    title: {
        fontSize: 18,
        marginBottom: 24,
    },
    button: {
        alignSelf: 'center',
    },
});
